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
	"os/exec"
	"time"
)

// Isosec returns a UTC timestamp (yyyymmddHHMMSS) as a string.
func isosec() string {
	t := time.Now().UTC()
	timeString := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return timeString
}

// EditFile opens file specified by filePath in editor.
func editFile(filePath string, editor string) error {
	cmd := exec.Command(editor, filePath)
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
