package fstraversal

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
)

const (
	IndentationSpace  = "    "
	CornerDelimiter   = "└── "
	VerticalDelimiter = "│   "
	CrossDelimiter    = "├── "
)

var dirStyle = lipgloss.NewStyle().Bold(true).Underline(true).Foreground(lipgloss.Color("5"))

func Dispatcher(options Options) {
	if options.Directory == "" {
		options.Directory = "."
	}

	if options.Flags.ShowTreeView {
		lines := []string{}
		lines = append(lines, dirStyle.Render(options.Directory))

		err := traverseFsTree(options.Directory, "", true, options.Flags.ShowHidden, &lines)

		if err != nil{
			fmt.Print(err)
			return
		}

		fmt.Print(lipgloss.JoinVertical(lipgloss.Left, lines...))
	} else {
		displayListDirectory(options.Directory, options.Flags.ShowHidden)
	}
}

func displayListDirectory(path string, showHidden bool) {
	files, err := os.ReadDir(path)

	if err != nil {
		fmt.Printf("Error reading directory: %s\n", err.Error())
		return
	}

	for _, file := range files {
		if !isHiddenFile(file.Name()) || showHidden {
			fmt.Println(formatOutputText("", file.Name(), file.IsDir()))
		}
	}
}

func traverseFsTree(path string, indent string, isLastFolder bool, showHidden bool, lines *[]string) error {
	files, err := os.ReadDir(path)

	if err != nil {
		return fmt.Errorf("error reading directory, %w", err)
	}

	lastFileIndx := getLastFileIndex(files, showHidden)

	for i, file := range files {
		if !isHiddenFile(file.Name()) || showHidden {
			var prefix, subDirectoryIndent string

			if (i == lastFileIndx && isLastFolder) || i == lastFileIndx {
				prefix = indent + CornerDelimiter
				subDirectoryIndent = indent + IndentationSpace
			} else {
				prefix = indent + CrossDelimiter
				subDirectoryIndent = indent + VerticalDelimiter
			}

			*lines = append(*lines, formatOutputText(prefix, file.Name(), file.IsDir()))

			if file.IsDir() {
				nextFilePath := filepath.Join(path, file.Name())
				newIsLastFolder := i == lastFileIndx && isLastFolder
				traverseFsTree(nextFilePath, subDirectoryIndent, newIsLastFolder, showHidden, lines)
			}
		}
	}

	return err
}

// Files prefixed with . are hidden from Unix ls command
func isHiddenFile(name string) bool {
	return len(name) > 0 && name[0] == '.'
}

func getLastFileIndex(files []fs.DirEntry, showHidden bool) int {
	for i := len(files) - 1; i >= 0; i-- {
		if !isHiddenFile(files[i].Name()) || showHidden {
			return i
		}
	}

	return -1
}

func formatOutputText(prefix string, fileName string, isDir bool) string {
	var format string

	if isDir {
		format = fmt.Sprintf("%s%s", prefix, dirStyle.Render(fileName))
	} else {
		format = fmt.Sprintf("%s%s", prefix, fileName)
	}
	return format
}
