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

// Package cmd implements subcommands and flags for the zet command.
package cmd

import (
	"os"
)

// Runner implements the requirements of a FlagSet runner type.
type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

// Root initializes all subcommands and flags, and runs the appropriate command.
func Root(args []string) error {
	if len(args) < 1 {
		return helpCmd().Run()
	}

	// all subcommands and flags are called here
	cmds := []Runner{
		newCmd(),
		helpCmd(),
		searchCmd(),
		editCmd(),
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
