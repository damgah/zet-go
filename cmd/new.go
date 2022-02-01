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
