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
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// SearchCommand encapsulates a *flag.FlagSet type. Flag names must be unique
// within a FlagSet. The FlagSet holds both subcommands and flags. The struct
// is also used to hold configuration related fields and search results.
type searchCommand struct {
	fs *flag.FlagSet

	in     string   // holds the -in flag
	zetdir string   // to store root directory
	titles []string // search results: md headers (titles)
	paths  []string // search results: absolute path to files
}

// SearchCmd returns a new searchCommand and implements the `search` subcommand
// and the `-in` flag. This is called by the Runner interface in root.go, which
// in turn calls the Name(), Init(), and Run() methods when called.
func searchCmd() *searchCommand {
	sc := &searchCommand{
		fs: flag.NewFlagSet("search", flag.ExitOnError),
	}

	sc.fs.StringVar(&sc.in, "in", "title", "must be either of {title, tags, all}")

	return sc
}

// Name returns the FlagSet method Name(). This is required by the Runner
// interface.
func (s *searchCommand) Name() string {
	return s.fs.Name()
}

// Init takes as input the arguments given by the user and parses these. It also
// initializes configuration fields.
func (s *searchCommand) Init(args []string) error {
	s.zetdir = os.Getenv("ZETDIR")
	if s.zetdir == "" {
		fmt.Println("configure ZETDIR environment variable and try again")
		os.Exit(1)
	}

	return s.fs.Parse(args)
}

// Run is called when the user calls the `search` subcommand. It calls the
// methods to perform the search and print the results.
func (s *searchCommand) Run() error {
	s.walkDirectory()
	s.printResults()
	return nil
}

// WalkDirectory recursively walks through the filesystem from the root of
// ZETDIR, calling searchFile on every README.md file. The headers and absolute
// paths of all files that match the search string at least once is recorded in
// the corresponding searchCommand field.
func (s *searchCommand) walkDirectory() error {
	// Walking the filesystem of ZETDIR from its root ("."), applying anonymous func to each file
	fs.WalkDir(os.DirFS(s.zetdir), ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		// ignore .git directory
		if d.IsDir() && d.Name() == ".git" {
			return filepath.SkipDir
		}

		// p is every file in ZETDIR (including directories)
		if strings.Contains(p, "README.md") {
			t, absPath := s.searchFile(p)

			if t != "" {
				d := strings.Split(p, "/")[0]
				s.titles = append(s.titles, d+"\t"+strings.Trim(t, "# "))
				s.paths = append(s.paths, absPath)
			}
		}
		return nil
	})
	return nil
}

// SearchFile searching titles, tags or entire files according to the in flag.
// It returns the absolute filepath to a file that contains at least one search
// string.
func (s *searchCommand) searchFile(p string) (string, string) {
	fp := s.zetdir + string(filepath.Separator) + p

	var (
		title   string
		absPath string
	)

	switch s.in {
	case "title":
		if len(s.fs.Args()) == 0 {
			fmt.Println("Enter search string")
			os.Exit(1)
		} else if s.fs.Args()[0] == "-in" || s.fs.Args()[0] == "--in" {
			if len(s.fs.Args()[2:]) == 0 {
				fmt.Println("Enter search string")
				os.Exit(1)
			} else {
				title, absPath = s.searchTitle(fp, s.fs.Args()[2:])
			}
		} else {
			title, absPath = s.searchTitle(fp, s.fs.Args())
		}
	case "tags":
		fmt.Println("Searching in:", s.in)
		fmt.Println("To be implemented")
	case "all":
		fmt.Println("Searching in:", s.in)
		fmt.Println("To be implemented")
	default:
		fmt.Println("search flag -in must be either of {title, tags, all}")
	}

	return title, absPath
}

// SearchTitle searches the first line of each file for strings that match args
func (s *searchCommand) searchTitle(fp string, args []string) (string, string) {
	// text := s.readFile(fp)
	// textLower := strings.ToLower(text[0]) // Title/search uses only first line of text
	text := s.readTitle(fp)
	textLower := strings.ToLower(text)

	var (
		match   string
		absPath string
	)

	for _, arg := range args {
		if strings.Contains(textLower, arg) {
			// match = text[0]
			match = text
			absPath = fp
		}
	}

	return match, absPath
}

// // ReadFile returns a slice of strings containing the content of the fp file
// func (s *searchCommand) readFile(fp string) []string {
// 	f, err := os.Open(fp)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	// Scan file line by line
// 	scanner := bufio.NewScanner(f)
// 	var text []string
// 	for scanner.Scan() {
// 		text = append(text, scanner.Text())
// 	}

// 	return text
// }

// ReadTitle returns a string containing the first line of the fp file
func (s *searchCommand) readTitle(fp string) string {
	f, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Scan first line of file
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text()
}

// PrintResults prints the search results
func (s *searchCommand) printResults() error {
	for i, value := range s.titles {
		fmt.Printf("[%d] %s\n", i, value)
	}
	return nil
}
