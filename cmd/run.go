package main

import (
	"fmt"

	"github.com/prtls/fstraversal"
	"github.com/spf13/cobra"
)

var flag fstraversal.CmdLineFlags

var RootCommand = &cobra.Command{
	Use:   "prtls",
	Short: "prtls is a CLI tool with custom styling for directory tree generation and ls commands.",
	Run:   runCobra,
}

func run() error {
	flagDefinition(&flag)
	if err := RootCommand.Execute(); err != nil {
		return err
	}

	return nil
}

func runCobra(_ *cobra.Command, args []string) {
	if validateDirectoryArgs(args) {
		fstraversal.Dispatcher(fstraversal.Options{Directory: args[len(args)-1], Flags: flag})
	} else if len(args) == 0 {
		fstraversal.Dispatcher(fstraversal.Options{Flags: flag})
	} else {
		fmt.Println("Usage: prtls [flags] <directory_path>")
	}
}

func flagDefinition(flags *fstraversal.CmdLineFlags) {
	flags.HideIcon = true
	RootCommand.PersistentFlags().BoolVarP(&flags.ShowHidden, "all", "a", false, "List all files (including hidden files), directories")
	RootCommand.PersistentFlags().BoolVarP(&flags.ShowTreeView, "tree", "t", false, "Tree view of the directory")
	RootCommand.PersistentFlags().BoolVarP(&flags.HideIcon, "icon", "i", false, "Disable icons (Enabled by default)")
}

func validateDirectoryArgs(args []string) bool {
	return len(args) == 1
}
