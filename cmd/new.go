package cmd

import (
	"flag"
	"fmt"
)

type newCommand struct {
	fs *flag.FlagSet
}

func newCmd() *newCommand {
	nc := &newCommand{
		fs: flag.NewFlagSet("new", flag.ExitOnError),
	}
	return nc
}

func (n *newCommand) Name() string {
	return n.fs.Name()
}

func (n *newCommand) Init(args []string) error {
	return n.fs.Parse(args)
}

func (n *newCommand) Run() error {
	fmt.Println("Ny kommando!")
	return nil
}
