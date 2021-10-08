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
	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/application"
)

type Deno struct {
	Javascript    *Javascript
	stages        []string
	DefinedStages []string
}

func NewDeno() Deno {
	return Deno{
		stages: []string{"lint", "format", "test"},
		Javascript: &Javascript{
			versions:          []string{"v1.x"},
			GithubActionsUser: "denoland/setup-deno@v1",
			GitLabBuildImage:  "",
			Language: &Language{
				Command: &Command{
					Linter:   []string{},
					Formater: []string{},
					Test:     []string{},
				},
			},
		},
	}
}

func (deno *Deno) Init(generate *application.Generate) {
	deno.defineStages()
	deno.Javascript.defineVersion()

	for _, stage := range deno.DefinedStages {
		if stage == "test" {
			deno.defineTest()
		}
		if stage == "lint" {
			deno.defineLint()
		}
		if stage == "format" {
			deno.defineFormat()
		}
	}

	deno.fill(generate)
}

func (deno *Deno) defineStages() {
	targetStages := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select the stages:",
		Help:    "Stages are the steps that the pipeline will cover.",
		Options: deno.stages[:],
	}
	survey.AskOne(prompt, &targetStages)
	deno.DefinedStages = targetStages
}

func (deno *Deno) defineLint() {
	deno.Javascript.Language.Command.Linter = append(deno.Javascript.Language.Command.Linter, "deno lint --unstable")
}

func (deno *Deno) defineFormat() {
	deno.Javascript.Language.Command.Formater = append(deno.Javascript.Language.Command.Formater, `deno fmt --check $(find . -iname "*.[j|t]s")`)
}

func (deno *Deno) defineTest() {
	deno.Javascript.Language.Command.Test = append(deno.Javascript.Language.Command.Test, `deno test --unstable --allow-all`)
}

func (deno *Deno) fill(generate *application.Generate) {
	generate.Language.Name = "Deno"
	generate.Language.Version = deno.Javascript.Version
	generate.DefinedBy = deno.DefinedStages
	generate.Language.Extension = []string{"js", "ts"}
	generate.Language.GitHubCiBuilder = deno.Javascript.GithubActionsUser
	generate.Language.GitLabCiBuilder = deno.Javascript.GitLabBuildImage

	generate.Command.Linter = deno.Javascript.Language.Command.Linter
	generate.Command.Formatter = deno.Javascript.Language.Command.Formater
	generate.Command.Test = deno.Javascript.Language.Command.Test
}
