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
	"log"
	"os"
	"strconv"
)

// EditCommand encapsulates a *flag.FlagSet type. Flag names must be unique
// within a FlagSet. The FlagSet holds both subcommands and flags.
type editCommand struct {
	fs *flag.FlagSet

	editor   string // to store editor
	repoUrl  string // to store url to repo
	zetTitle string // to store title
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
	e.editor = os.Getenv("EDITOR")
	if e.editor == "" {
		fmt.Println("configure EDITOR environment variable and try again")
		os.Exit(1)
	}

	return e.fs.Parse(args)
}

// Run is called when the user calls the `edit` subcommand.
func (e *editCommand) Run() error {
	switch len(e.fs.Args()) {
	case 0:
		fmt.Println("Enter search string to search titles")
	default:
		// perform search and store results
		s := searchCmd()
		s.Init(e.fs.Args())
		s.Run()

		// scan user input as string
		fmt.Println("Edit?")
		var v string
		fmt.Scan(&v)

		// exit program if `v` is not convertible to int
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}

		// edit file
		if i >= 0 && i < len(s.titles) {
			e.zetTitle = s.titles[i]
			err = editFile(s.paths[i], e.editor)
			if err != nil {
				log.Fatal(err)
			}
		}

		// Get remote repo url
		e.repoUrl, err = getRemoteGitUrl(s.zetdir)
		if err != nil {
			log.Fatal(err)
		}

		// Add and commit changes
		addAndCommit(s.zetdir, e.zetTitle)

		// Pull and push changes
		pullFromRepo(s.zetdir, e.repoUrl)
		pushToRepo(s.zetdir, e.repoUrl)
	}

	return nil
}
