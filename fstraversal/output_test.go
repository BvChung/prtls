package fstraversal

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
)

func TestRemovingANSI(t *testing.T) {
	assert := assert.New(t)

	testData := []struct {
		input    string
		expected string
	}{
		{
			input:    "\033[31mThis is a red text\033[0m",
			expected: "This is a red text",
		},
		{
			input:    "\033[32mThis is a green text\033[0m",
			expected: "This is a green text",
		},
		{
			input:    "\033[33mThis is a yellow text\033[0m",
			expected: "This is a yellow text",
		},
		{
			input:    "\033[34mThis is a blue text\033[0m",
			expected: "This is a blue text",
		},
		{
			input:    "\033[35mThis is a magenta text\033[0m",
			expected: "This is a magenta text",
		},
		{
			input:    "\033[36mThis is a cyan text\033[0m",
			expected: "This is a cyan text",
		},
	}

	for _, d := range testData {
		t.Run("Test remove ANSI", func(t *testing.T) {
			removedANSIStr := removeANSI(d.input)
			assert.Equal(removedANSIStr, d.expected)
		})
	}
}

func TestTreeCleaning(t *testing.T) {
	assert := assert.New(t)

	testData := []struct {
		expected []string
		path     string
	}{
		{
			expected: []string{
				"└── " + dirStyle.Render("b"),
				"    └── " + dirStyle.Render("c"),
				"        ├── cc.py",
				"        └── " + dirStyle.Render("d"),
				"            ├── a.txt",
				"            └── b.txt",
			},
			path: "../testdata/mockDirectory"},
	}

	for _, d := range testData {
		t.Run("Test tree cleaning", func(t *testing.T) {
			lines := []string{}
			if err := traverseFsTree(d.path, "", false, &lines); err != nil {
				t.Fatal(err)
			}

			expectedTree := lipgloss.JoinVertical(lipgloss.Left, d.expected...)
			actualTree := lipgloss.JoinVertical(lipgloss.Left, lines...)

			expectedCleanTree := cleanTree(expectedTree)
			actualTreeCleanTree := cleanTree(actualTree)

			assert.Equal(expectedCleanTree, actualTreeCleanTree)
		})
	}
}
