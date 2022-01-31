package cmd

import (
	"flag"
	"fmt"
)

type helpCommand struct {
	fs *flag.FlagSet
}

func helpCmd() *helpCommand {
	hc := &helpCommand{
		fs: flag.NewFlagSet("help", flag.ExitOnError),
	}
	return hc
}

func (h *helpCommand) Name() string {
	return h.fs.Name()
}

func (h *helpCommand) Init(args []string) error {
	return h.fs.Parse(args)
}

func (h *helpCommand) Run() error {
	fmt.Println("Usage:")
	fmt.Println("zet {cmd} [flags]")
	fmt.Println("where {cmd} is one of")
	fmt.Println("\thelp: displays this usage information")
	fmt.Println("\tnew: create a new zet direcotory and file, and start editor")
	fmt.Println("\tsearch [-title, -tags, -all] {string}: search zet repo for {string} in title (default), tags, or entire file")
	fmt.Println("\tedit {string}: numbered list of zets containing {string} in title, and opens selected file in editor")
	fmt.Println("\tcommit: commits changes using header of added file as commit message, pulls, and pushes to github")
	return nil
}
