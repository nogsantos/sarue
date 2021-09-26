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
	"github.com/nogsantos/sarue/utils"
)

type Node struct {
	Javascript *Javascript
	// Setup
	managers []string
	platforms []string
	frontFrameworks []string
	backendFrameworks []string
	stages []string
	// Parameters
	Platform string
	Manager string
	FrontFramework string
	BackendFramework string
	DefinedStages []string
}

func NewNode() Node {
	return Node{
		stages: []string{"lint", "format", "test"},
		managers: []string{"npm", "yarn 1", "yarn 2", "pnpm", "None"},
		platforms: []string{"Fronted", "Backend"},
		frontFrameworks: []string{"Vue", "React", "None"},
		backendFrameworks: []string{"Nest", "Vuex", "Next"},
		Javascript: &Javascript{
			versions: []string{"15.x", "14.x", "12.x", "10.x"},
			GithubActionsUser: "actions/setup-node@v2",
			GitLabBuildImage: "",
			Language: &Language{
				Command: &Command{
					Linter: []string{},
					Formater: []string{},
					Test: []string{},
				},
			},
		},
	}
}

func (node *Node) Init(generate *application.Generate) {
	node.defineStages()
	node.definePlatform()
	node.defineManager()
	// node.defineFrameworks()

	node.Javascript.defineVersion()

	for _, stage := range node.DefinedStages {
		if stage == "test" {
			node.defineTest()
		}
		if stage == "lint" {
			node.defineLint()
		}
		if stage == "format" {
			node.defineFormat()
		}
	}

	node.fill(generate)
}

func (node *Node) defineStages() {
	targetStages := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select the stages:",
		Help: "Stages are the steps that the pipeline will cover.",
		Options: node.stages[:],

	}
	survey.AskOne(prompt, &targetStages)
	node.DefinedStages = targetStages
}

func (node *Node) defineManager() {
	targetManager := ""
	prompt := &survey.Select{
		Message: "Using a manager?",
		Options: node.managers[:],
	}
	err := survey.AskOne(prompt, &targetManager)
	if err != nil {
		utils.Error(err.Error())
	}
	node.Manager = targetManager
}

func (node *Node) definePlatform() {
	targetPlatform := ""
	platformPrompt := &survey.Select{
		Message: "Will be used for front or backend:",
		Options: node.platforms[:],
	}
	err := survey.AskOne(platformPrompt, &targetPlatform)
	if err != nil {
		utils.Error(err.Error())
	}
	node.Platform = targetPlatform

	switch targetPlatform {
		case node.platforms[0]:
			node.frontendHandler()
		case node.platforms[1]:
			node.backendHandler()
	default:
		utils.Error("Node")
	}
}

func (node *Node) frontendHandler() {
	targetFramework := ""
	prompt := &survey.Select{
		Message: "What is the framework:",
		Options: node.frontFrameworks[:],
	}
	err := survey.AskOne(prompt, &targetFramework)
	if err != nil {
		utils.Error(err.Error())
	}
	node.BackendFramework = ""
	node.FrontFramework = targetFramework
}

func (node *Node) backendHandler() {
	targetFramework := ""
	prompt := &survey.Select{
		Message: "What is the framework:",
		Options: node.backendFrameworks[:],
	}
	err := survey.AskOne(prompt, &targetFramework)
	if err != nil {
		utils.Error(err.Error())
	}
	node.FrontFramework = ""
	node.BackendFramework = targetFramework
}

func (node *Node) defineLint() {
	node.Javascript.Language.Command.Linter = append(
		node.Javascript.Language.Command.Linter,
		node.Manager + " eslint --ext .js --ext .ts .",
	)
}

func (node *Node) defineFormat() {
	node.Javascript.Language.Command.Formater = append(
		node.Javascript.Language.Command.Formater,
		node.Manager + ` prettier --no-error-on-unmatched-pattern --check "**/*.js" "**/*.ts"`,
	)
}

func (node *Node) defineTest() {
	node.Javascript.Language.Command.Test = append(
		node.Javascript.Language.Command.Test,
		node.Manager + ` run test`,
	)
}

func (node *Node) fill(generate *application.Generate) {
	generate.Language.Name = "Node"
	generate.ManagerName = node.Manager
	if node.FrontFramework != "" {
		generate.FrameworkName = node.FrontFramework
	} else {
		generate.FrameworkName = node.BackendFramework
	}
	generate.Language.Version = node.Javascript.Version
	generate.DefinedBy = node.DefinedStages
	generate.Language.Extension = []string{"js", "ts"}
	generate.Language.GitHubCiBuilder = node.Javascript.GithubActionsUser
	generate.Language.GitLabCiBuilder = node.Javascript.GitLabBuildImage

	generate.Command.Linter = node.Javascript.Language.Command.Linter
	generate.Command.Formatter = node.Javascript.Language.Command.Formater
	generate.Command.Test = node.Javascript.Language.Command.Test
}