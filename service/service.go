package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
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
	fmt.Println(formatText("."))
	if options.Flags.ShowTreeView {
		displayTreeDirectory(options, "", true)
	} else {
		displayListDirectory(options)
	}
}

func displayListDirectory(options Options) {
	CheckDefaultDirectory(&options.Directory)

	files, dir, err := readDirectory(options.Directory)
	defer closeDirectory(dir)

	if err != nil {
		fmt.Printf("Error reading directory: %s\n", err.Error())
		return
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if !isHidden(file.Name()) || options.Flags.ShowHidden {
			fmt.Println(formatFileDirectories("", file))
		}
	}
}

func displayTreeDirectory(options Options, indent string, isLastFolder bool) {
	CheckDefaultDirectory(&options.Directory)

	files, dir, err := readDirectory(options.Directory)
	defer closeDirectory(dir)

	if err != nil {
		log.Fatalf("Error reading directory: %s\n", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	lastVisibleIndex := getLastVisibleIndex(files, options.Flags.ShowHidden)
	for i, file := range files {
		if !isHidden(file.Name()) || options.Flags.ShowHidden {
			var prefix, childIndentNext string

			if i == lastVisibleIndex && isLastFolder {
				prefix = indent + CornerDelimiter
				childIndentNext = indent + IndentationSpace
			} else if i == lastVisibleIndex {
				prefix = indent + CornerDelimiter
				childIndentNext = indent + IndentationSpace
			} else {
				prefix = indent + CrossDelimiter
				childIndentNext = indent + VerticalDelimiter
			}

			fmt.Println(formatFileDirectories(prefix, file))

			if file.IsDir() {
				newDirectory := filepath.Join(options.Directory, file.Name())
				newIsLastFolder := i == lastVisibleIndex && isLastFolder
				displayTreeDirectory(Options{Directory: newDirectory, Flags: CmdLineFlags{ShowHidden: options.Flags.ShowHidden, HideIcon: options.Flags.HideIcon}}, childIndentNext, newIsLastFolder)
			}
		}
	}
}

func getLastVisibleIndex(files []os.FileInfo, showHidden bool) int {
	for i := len(files) - 1; i >= 0; i-- {
		if !isHidden(files[i].Name()) || showHidden {
			return i
		}
	}
	return -1
}
