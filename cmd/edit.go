/*
Copyright 2022 Damoun Ashournia.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"flag"
	"fmt"
)

// EditCommand encapsulates a *flag.FlagSet type. Flag names must be unique
// within a FlagSet. The FlagSet holds both subcommands and flags.
type editCommand struct {
	fs *flag.FlagSet
}

// EditCmd returns a new editCommand and implements the `edit` subcommand.
// This is called by the Runner interface in root.go, which in turn calls the
// Name(), Init(), and Run() methods when called.
func editCmd() *editCommand {
	ec := &editCommand{
		fs: flag.NewFlagSet("edit", flag.ExitOnError),
	}
	return ec
}

// Name returns the FlagSet method Name(). This is required by the Runner
// interface.
func (e *editCommand) Name() string {
	return e.fs.Name()
}

// Init takes as input the arguments given by the user and parses these.
func (e *editCommand) Init(args []string) error {
	return e.fs.Parse(args)
}

// Run is called when the user calls the `edit` subcommand.
func (e *editCommand) Run() error {
	// perform search and store results
	switch len(e.fs.Args()) {
	case 0:
		fmt.Println("Enter search string to search titles")
	default:
		s := searchCmd()
		s.Init(e.fs.Args())
		s.Run()
		// TODO: Scan user input as int, check if valid comparing to number of
		// search results, if valid open in editor
	}

	return nil
}
