package cmd

import (
	"os"
)

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func Root(args []string) error {
	if len(args) < 1 {
		return helpCmd().Run()
	}

	cmds := []Runner{
		newCmd(),
		helpCmd(),
		searchCmd(),
	}

	subcommand := os.Args[1]

	for _, c := range cmds {
		if c.Name() == subcommand {
			c.Init(os.Args[2:])
			return c.Run()
		}
	}

	return helpCmd().Run()
}
