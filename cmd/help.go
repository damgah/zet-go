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
	fmt.Println("\tnew: create a new zettel directory and file, and start editor")
	fmt.Println("\tsearch [-in (title, tags, all)] {string}: search repo for {string} in title, tags, or entire file (default)")
	fmt.Println("\tedit {string}: numbered list of zettels containing {string} in file, and open selected file in editor")
	return nil
}
