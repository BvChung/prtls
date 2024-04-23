package fstraversal

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type model struct {
	filepicker     filepicker.Model
	fpActive       bool
	showTree       bool
	treeStructure  string
	outputFilePath string

	viewport   viewport.Model
	viewActive bool

	currentDir  string
	selectedDir string

	content  string
	ready    bool
	quitting bool
}

var (
	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("212")).
			Bold(true).
			Padding(0, 1)

	sepStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("99"))
	helperStyle = lipgloss.NewStyle().Bold(true)
	dirStyle    = lipgloss.NewStyle().Bold(true).Underline(true).Foreground(lipgloss.Color("99"))
	rootStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212"))
)

func NewModel(initPath string, outputFilePath string) model {
	fp := filepicker.New()
	fp.DirAllowed = true
	fp.FileAllowed = false
	fp.CurrentDirectory = initPath

	return model{
		filepicker:     fp,
		fpActive:       true,
		viewActive:     false,
		currentDir:     initPath,
		selectedDir:    initPath,
		outputFilePath: outputFilePath,
	}
}

func (m model) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	if m.fpActive {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "t":
				m.fpActive = false
				m.showTree = true
				m.viewActive = true
				m.treeStructure = createTree(m.currentDir, m.filepicker.ShowHidden)
				m.viewport.SetContent(m.treeStructure)
			case "backspace", "b":
				m.showTree = false
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			}

		case tea.WindowSizeMsg:
			fpHeaderHeight := lipgloss.Height(m.headerView())
			fpFooterHeight := lipgloss.Height(m.fpFooterView())
			m.filepicker.Height = msg.Height - (fpHeaderHeight + fpFooterHeight)

			viewHeaderHeight := lipgloss.Height(m.headerView())
			viewFooterHeight := lipgloss.Height(m.viewFooterView())
			verticalMarginHeight := viewHeaderHeight + viewFooterHeight

			if !m.ready {
				m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
				m.viewport.YPosition = viewHeaderHeight
				m.viewport.HighPerformanceRendering = false

				m.viewport.SetContent(m.content)
				m.ready = true

				m.viewport.YPosition = viewHeaderHeight + 1
			} else {
				m.viewport.Width = msg.Width
				m.viewport.Height = msg.Height - verticalMarginHeight
			}
		}

		m.filepicker, cmd = m.filepicker.Update(msg)
		cmds = append(cmds, cmd)
		m.currentDir = m.filepicker.CurrentDirectory

		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.viewFooterView())
		verticalMarginHeight := headerHeight + footerHeight

		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - verticalMarginHeight
	case tea.KeyMsg:
		switch msg.String() {
		case "s":
			cleanedTree := cleanTree(m.treeStructure)
			if err := writeToFile(m.outputFilePath, []byte(cleanedTree)); err != nil {
				log.Error(err)
				return m, tea.Quit
			}
		case "backspace", "b":
			m.showTree = false
			m.fpActive = true
			m.viewActive = false
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)

}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	if m.viewActive && m.showTree {
		return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.viewFooterView())
	} else {
		return fmt.Sprintf("%s\n%s%s", m.headerView(), m.filepicker.View(), m.fpFooterView())
	}

}

func (m model) headerView() string {
	return infoStyle.Render("Current Directory: " + m.filepicker.Styles.Selected.Render(m.currentDir))
}

func (m model) viewFooterView() string {
	cmds := []string{"↑/↓/scroll move", "s save", "b/backspace return", "q/ctrl+c quit"}
	s := infoStyle.Render(buildHelperCmdString(cmds))
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)-lipgloss.Width(s)))
	return lipgloss.JoinHorizontal(lipgloss.Center, s, line, info)
}

func (m model) fpFooterView() string {
	cmds := []string{"↑/↓/←/→ move", "t show dir hierarchy", "q/ctrl+c quit"}

	return helperStyle.Render(buildHelperCmdString(cmds))
}

func buildHelperCmdString(cmds []string) string {
	var b strings.Builder

	for idx, s := range cmds {
		if idx == len(cmds)-1 {
			b.WriteString(helperStyle.Render(s))
		} else {
			b.WriteString(fmt.Sprintf("%s %s ", helperStyle.Render(s), sepStyle.Render("•")))
		}
	}

	return b.String()
}