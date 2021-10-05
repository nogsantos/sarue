/*
Copyright 2021 The Sarue Authors.

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

package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/utils"
)

type Git struct {}

// Git confirms if the repository has a valid git configuration
func (git *Git) Init() {
	_, readFileError := ioutil.ReadFile(".git/config")

	if readFileError != nil {
		git.initRepo()
	}

	// hasRemoteEntry := strings.Contains(string(content), "remote \"origin\"")

	// if hasRemoteEntry {
	// 	fmt.Println("hasRemoteEntry")
	// } else {
	// 	fmt.Println("dont hasRemoteEntry")
	// }
}

// init Initialize a empty repo by user confirmation
func (git *Git) initRepo() {
	initRepo := false
	prompt := &survey.Confirm{
		Message: "No git configuration file was found. Do you want us to initialize it for you?",
		Default: true,
	}
	err := survey.AskOne(prompt, &initRepo)
	if err != nil {
		utils.Error(err.Error())
	}

	if !initRepo {
		fmt.Println("To create the configuration you must initialize your repo before.")
		os.Exit(0)
	}

	fmt.Println("Initializing the repo...")
	cmd := exec.Command("git", "init")
    stdout, commandError := cmd.Output()
	if commandError != nil {
		fmt.Printf("Git error %v", commandError)
		os.Exit(0)
	}
	fmt.Printf("%v", string(stdout))
}
