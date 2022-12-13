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
	"path/filepath"
	"strings"
	"text/template"
)

// NewCommand encapsulates a *flag.FlagSet type. Flag names must be unique
// within a FlagSet. The FlagSet may hold both subcommands and flags. The
// struct is also used to hold configuration related fields.
type newCommand struct {
	fs *flag.FlagSet

	zetdir          string // to store root directory
	editor          string // to store editor
	timeString      string // to store the timestamp
	timeStringIndex string // to store the timestamp for generating the index
}

// NewCmd creates a new zettelkasten file with the header given by the user
// passed arguments. The zettelkasten is placed in a new directory with isosec
// name. This is called by the Runner interface in root.go, which in turn calls
// the Name(), Init(), and Run() methods when called.
func newCmd() *newCommand {
	nc := &newCommand{
		fs: flag.NewFlagSet("new", flag.ExitOnError),
	}
	return nc
}

// Name returns the FlagSet method Name(). This is required by the Runner
// interface.
func (n *newCommand) Name() string {
	return n.fs.Name()
}

// Init takes as input the arguments given by the user and parses these. It also
// initializes configuration fields.
func (n *newCommand) Init(args []string) error {
	n.zetdir = os.Getenv("ZETDIR")
	if n.zetdir == "" {
		fmt.Println("configure ZETDIR environment variable and try again")
		os.Exit(1)
	}

	n.editor = os.Getenv("EDITOR")
	if n.editor == "" {
		fmt.Println("configure EDITOR environment variable and try again")
		os.Exit(1)
	}

	return n.fs.Parse(args)
}

// Run is called when the user calls the `new` subcommand. It calls the
// methods to perform the folder and file creation, and opens the editor.
func (n *newCommand) Run() error {
	// Set timestamps
	n.timeString, n.timeStringIndex = isosec()

	filePath := n.newFile()

	err := editFile(filePath, n.editor)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// NewFile creates a file named README.md in the appropriate directory. The
// file is created from a template that uses optional arguments passed by the
// user to create a header.
func (n *newCommand) newFile() string {
	// Create directory
	dirPath := n.createDirectory()

	// Create file
	filePath := dirPath + string(filepath.Separator) + "README.md"
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get header from arguments passed by user
	type header struct {
		Title string
	}
	h := header{strings.Join(n.fs.Args(), " ")}

	// Create template
	f := template.Must(template.New("temp").Parse("# {{.Title}}"))

	// Write template to file
	err = f.Execute(file, h)
	if err != nil {
		log.Fatal(err)
	}

	// Update index file
	indexPath := n.zetdir + string(filepath.Separator) + "README.md"
	line := "* " + n.timeStringIndex + " [" + h.Title + "](" + "https://github.com/damgah/zet/tree/main/zets/" + n.timeString + "/README.md)"
	UpdateIndexFile(indexPath, line)

	return filePath
}

// CreateDirectory creates a directory for new file named by the current UTC
// time. It returns the path to that directory.
func (n *newCommand) createDirectory() string {
	path := n.zetdir + string(filepath.Separator) + "zets" + string(filepath.Separator) + n.timeString
	os.Mkdir(path, 0755)
	return path
}

// UpdateIndexFile appends a line of text to the existing index file.
// If the file doesn't exist, it creates a new file before appending the line.
func UpdateIndexFile(filePath string, line string) error {
	// Open the file in append mode
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// The file doesn't exist, so we create a new file
		file, err = os.Create(filePath)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	// Append the line of text to the file
	_, err = file.WriteString(line + "\n")
	if err != nil {
		return err
	}

	// The line has been appended successfully
	return nil
}
