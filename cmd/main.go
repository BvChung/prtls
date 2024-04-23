package main

import (
	"flag"
	"fmt"
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
	flag.StringVar(&initPath, "p", ".", "-p | -p=")
	flag.StringVar(&outputPath, "o", "", "-o | -o=")
	flag.Parse()

	fmt.Println(outputPath)

	m := fstraversal.NewModel(initPath, outputPath)

	teaProgram := tea.NewProgram(&m, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := teaProgram.Run(); err != nil {
		return err
	}

	return nil
}
