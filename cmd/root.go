package cmd

import (
	"fmt"
	"log"

	"github.com/prtls/service"
	"github.com/spf13/cobra"
)

var flag service.CmdLineFlags

var RootCommand = &cobra.Command{
	Use:   "prtls",
	Short: "prtls is a CLI tool with custom styling for tree and ls commands.",
	Run:   run,
}

func Execute() {
	flagDefinition(&flag)
	if err := RootCommand.Execute(); err != nil {
		log.Fatalf(err.Error())
	}
}

func run(_ *cobra.Command, args []string) {
	if validateDirectoryArgs(args) {
		service.Dispatcher(service.Options{Directory: args[len(args)-1], Flags: flag})
	} else if len(args) == 0 {
		service.Dispatcher(service.Options{Flags: flag})
	} else {
		fmt.Println("Usage: prtls [options] <directory_path>")
	}
}

func flagDefinition(flags *service.CmdLineFlags) {
	flags.HideIcon = true
	RootCommand.PersistentFlags().BoolVarP(&flags.ShowHidden, "all", "a", false, "List all files (including hidden files), directories")
	RootCommand.PersistentFlags().BoolVarP(&flags.ShowTreeView, "tree", "t", false, "Tree view of the directory")
	RootCommand.PersistentFlags().BoolVarP(&flags.HideIcon, "icon", "i", false, "Disable icons (Enabled by default)")
}

func validateDirectoryArgs(args []string) bool {
	return len(args) == 1
}
