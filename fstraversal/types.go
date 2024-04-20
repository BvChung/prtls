package fstraversal

type CmdLineFlags struct {
	ShowHidden   bool
	ShowTreeView bool
	HideIcon     bool
}

type Options struct {
	Directory string
	Flags     CmdLineFlags
}

