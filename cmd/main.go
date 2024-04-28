package main

import (
	"flag"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/prtls/fstraversal"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	var initPath string
	var outputPath string
	flag.StringVar(&initPath, "p", ".", "-p, -p=, --p string	optional initial directory path (defaults to current directory)")
	flag.StringVar(&outputPath, "o", "", "-o, -o=, --o string	output file path")
	flag.Parse()

	m := fstraversal.NewModel(initPath, outputPath)

	teaProgram := tea.NewProgram(&m, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := teaProgram.Run(); err != nil {
		return err
	}

	return nil
}
