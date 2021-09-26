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

package language

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/utils"
)

type Javascript struct {
	Language *Language
	// Setup
	versions []string
	// Parameters
	Runtime string
	Framework string
	Version string

	GithubActionsUser string
	GitLabBuildImage string
}

func NewJavascript(runtime string) *Javascript {
	return &Javascript{
		Runtime: runtime,
	}
}

func (js *Javascript) Init(generate *application.Generate) {
	fmt.Println("JS INIT")
	if js.Runtime == NODE.String() {
		nd := NewNode()
		nd.Init(generate)
	} else {
		fmt.Println("DENO INIT")
		dn := NewDeno()
		dn.Init(generate)
	}
}

func (js *Javascript) defineVersion() {
	fmt.Print("VERSION")
	targetVersion := ""
	prompt := &survey.Select{
		Message: "What is the runtime version?",
		Options: js.versions[:],
	}
	err := survey.AskOne(prompt, &targetVersion)
	if err != nil {
		utils.Error(err.Error())
	}
	js.Version = targetVersion
}

func (js *Javascript) defineManager() {}
func (js *Javascript) defineFramework() {}
func (js *Javascript) defineLint() {}
func (js *Javascript) defineTest() {}
func (js *Javascript) defineFormat() {}
func (js *Javascript) fill(generate *application.Generate) {}