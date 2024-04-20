package fstraversal

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/prtls/assets"
)

type CmdLineFlags struct {
	ShowHidden   bool
	ShowTreeView bool
	HideIcon     bool
}

type Options struct {
	Directory string
	Flags     CmdLineFlags
}

const (
	IndentationSpace  = "    "
	CornerDelimiter   = "└── "
	VerticalDelimiter = "│   "
	CrossDelimiter    = "├── "
)

func Dispatcher(options Options) {
	if options.Directory == "" {
		options.Directory = "."
	}

	fmt.Println(formatText(options.Directory))
	if options.Flags.ShowTreeView {
		displayTreeDirectory(options.Directory, "", true, options.Flags.ShowHidden)
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
			fmt.Println(formatFileOutput("", file))
		}
	}
}

func displayTreeDirectory(path string, indent string, isLastFolder bool, showHidden bool) {
	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatalf("Error reading directory: %s\n", err)
	}

	lastFileIndx := getLastFileIndex(files, showHidden)
	
	for i, file := range files {
		if !isHiddenFile(file.Name()) || showHidden {
			var prefix, subDirectoryIndent string

			if i == lastFileIndx && isLastFolder || i == lastFileIndx {
				prefix = indent + CornerDelimiter
				subDirectoryIndent = indent + IndentationSpace
			} else {
				prefix = indent + CrossDelimiter
				subDirectoryIndent = indent + VerticalDelimiter
			}

			fmt.Println(formatFileOutput(prefix, file))

			if file.IsDir() {
				nextFilePath := filepath.Join(path, file.Name())
				newIsLastFolder := i == lastFileIndx && isLastFolder
				displayTreeDirectory(nextFilePath, subDirectoryIndent, newIsLastFolder, showHidden)
			}
		}
	}
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

func formatFileOutput(prefix string, file fs.DirEntry) string {
	var format string
	if file.IsDir() {
		format = fmt.Sprintf("%s%s%s%s%s", prefix, assets.Bold, assets.Pink, file.Name(), assets.Reset)
	} else {
		format = fmt.Sprintf("%s%s", prefix, file.Name())
	}
	return format
}

func formatText(str string) string {
	return fmt.Sprintf("%s%s%s%s", assets.Bold, assets.LightBlue, str, assets.Reset)
}