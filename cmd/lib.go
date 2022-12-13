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
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Isosec returns a UTC timestamp (yyyymmddHHMMSS) as a string.
func isosec() (string, string) {
	t := time.Now().UTC()
	timeString := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	timeStringIndex := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return timeString, timeStringIndex
}

// EditFile opens file specified by filePath in editor.
func editFile(filePath string, editor string) error {
	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Start()
	if err != nil {
		fmt.Println("failed to launch editor")
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("finished with error:", err)
	}
	return nil
}

func addAndCommit(localRepoPath string, commitMessage string) {
	// Use the 'git' command to add and commit the changes in the local repo
	addCmd := exec.Command("git", "--git-dir", localRepoPath+"/.git", "--work-tree", localRepoPath, "add", ".")
	addCmd.Run()

	commitCmd := exec.Command("git", "--git-dir", localRepoPath+"/.git", "--work-tree", localRepoPath, "commit", "-m", commitMessage)
	commitCmd.Run()

	fmt.Println("Changes committed to local repo at", localRepoPath)
}

func pullFromRepo(localRepoPath string, repoUrl string) {
	cmd := exec.Command("git", "--git-dir", localRepoPath+"/.git", "--work-tree", localRepoPath, "pull", repoUrl)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error pulling from repository: ", err)
	} else {
		fmt.Println(string(output))
	}
}

func pushToRepo(localRepoPath string, repoUrl string) {
	cmd := exec.Command("git", "--git-dir", localRepoPath+"/.git", "--work-tree", localRepoPath, "push", repoUrl)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error pushing to repository: ", err)
	} else {
		fmt.Println(string(output))
	}
}

func getRemoteGitUrl(localRepoPath string) (string, error) {
	cmd := exec.Command("git", "--git-dir", localRepoPath+"/.git", "--work-tree", localRepoPath, "config", "--get", "remote.origin.url")
	output, err := cmd.Output()

	if err != nil {
		return "", err
	} else {
		return strings.TrimSpace(string(output)), nil
	}
}
