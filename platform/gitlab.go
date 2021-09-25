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

package platform

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/utils"
)

type GitLab struct {
	ConfigFile string
	ConfigPath string

}

func (gl *GitLab) construct() {
	gl.ConfigFile = ".gitlab-ci.yml"
	gl.ConfigPath = "./"
}

func (gl *GitLab) Init(generate *application.Generate) {
	gl.construct()
	gl.gitlabConfigExists(generate)
	gl.configureGitlaCi(generate)
}

func (gl *GitLab) gitlabConfigExists(generate *application.Generate) {
	path := fmt.Sprintf("%v/%v", utils.PathRoot(), gl.ConfigFile)
	_, readFileError := ioutil.ReadFile(path)
	if readFileError == nil {
		gl.confirmReplacement(generate)
	}
}

func (gl *GitLab) confirmReplacement(generate *application.Generate) {
	relplaceConfig := false
	prompt := &survey.Confirm{
		Message: "A Gitlab CI configuration was found, do you want to replace it?",
		Default: false,
	}
	err := survey.AskOne(prompt, &relplaceConfig)
	if err != nil {
		utils.Error(err.Error())
	}

	if !relplaceConfig {
		fmt.Printf("Finishing the process. We have nothing to do here.")
		os.Exit(0)
	} else {
		gl.configureGitlaCi(generate)
	}
}

func (gl *GitLab) configureGitlaCi(generate *application.Generate) {
	utils.WriteConfigFile(gl.ConfigPath, gl.ConfigFile, gl.GitlabTemplate(), generate)
}
