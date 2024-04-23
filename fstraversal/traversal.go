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

func createTree(directory string, showHidden bool, root string) string {
	lines := []string{}
	lines = append(lines, rootStyle.Render(root))
	traverseFsTree(directory, "", true, showHidden, &lines)

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
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
