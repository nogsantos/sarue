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
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/utils"
)

type Node struct {
	Javascript *Javascript
	managers []string
	stages []string
	Platform string
	Manager string
	FrontFramework string
	BackendFramework string
	DefinedStages []string
	PackageJson PackageJson
}

type PackageJson struct {
    Scripts struct {
		Build string `json:"build"`
		Format string `json:"format"`
		Test string `json:"test"`
		TestUnit string `json:"test:unit"`
		Lint string `json:"lint"`
	} `json:"scripts"`
}

func NewNode() Node {
	return Node{
		stages: []string{"lint", "format", "test"},
		managers: []string{"npm", "yarn 1", "yarn 2", "pnpm", "None"},
		Javascript: &Javascript{
			versions: []string{"16.x", "14.x", "12.x"},
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
	node.defineManager()
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

	err = node.loadConfiguration()
	if err != nil {
		log.Println(err.Error(), "Using npx by default")
	}

	hasConfigurationFile := (
		node.PackageJson.Scripts.Build != "" ||
		node.PackageJson.Scripts.Format != "" ||
		node.PackageJson.Scripts.Lint != "" ||
		node.PackageJson.Scripts.Test != "" ||
		node.PackageJson.Scripts.TestUnit != "")

	if hasConfigurationFile {
		// Install the manager Install the packages by manager definition
		switch node.Manager {
			case "npm":
				node.setupNpm()
			case "yarn 1":
				node.setupYarn()
			case "yarn 2":
				node.setupYarnTwo()
			case "pnpm":
				node.setupPnpm()
			default:
				node.setupNpm()
		}
	} else {
		log.Println("no script definition was found in the package.json file. Using npx by default")
	}
}

// setupNpm defines the npm setup for the project
// NPM: https://docs.npmjs.com/cli/v7/commands/npm-ci
func (node *Node) setupNpm() {
	node.Javascript.Language.Command.Build = append(
		node.Javascript.Language.Command.Build,
		"npm ci",
	)
}

// setupYarn defines yarn setup to the project installing its dependencies
// YARN 1: https://classic.yarnpkg.com/en/docs/cli/
func (node *Node) setupYarn() {
	node.Javascript.Language.Command.Build = append(
		node.Javascript.Language.Command.Build,
		"npm install --global yarn",
	)
	node.Javascript.Language.Command.Build = append(
		node.Javascript.Language.Command.Build,
		"yarn install --non-interactive",
	)
}

// setupYarn defines yarn 2 setup to the project installing its dependencies
// YARN 2: https://yarnpkg.com/cli/
func (node *Node) setupYarnTwo() {
	node.Javascript.Language.Command.Build = append(
		node.Javascript.Language.Command.Build,
		"npm install -g yarn",
	)
	node.Javascript.Language.Command.Build = append(
		node.Javascript.Language.Command.Build,
		"yarn set version from sources",
	)
	node.Javascript.Language.Command.Build = append(
		node.Javascript.Language.Command.Build,
		"yarn install --immutable --immutable-cache --check-cache",
	)
}

// setupYarn defines pnpm setup to the project installing its dependencies
// PNPM: https://pnpm.io/continuous-integration
func (node *Node) setupPnpm() {
	node.Javascript.Language.Command.Build = append(
		node.Javascript.Language.Command.Build,
		"npm install -g pnpm",
	)
	node.Javascript.Language.Command.Build = append(
		node.Javascript.Language.Command.Build,
		"pnpm add -g pnpm",
	)
	node.Javascript.Language.Command.Build = append(
		node.Javascript.Language.Command.Build,
		"pnpm install",
	)
}

// loadConfiguration read a node configuration file and
// set the enabled scripts in PackageJson structure
func (node *Node) loadConfiguration() error {
	packageJson := PackageJson{}
	packageJsonFile, err := os.Open("package.json")
	if err != nil {
		return errors.New("a package.json file was not found in the project root directory")
	}
	defer packageJsonFile.Close()

	jsonParser := json.NewDecoder(packageJsonFile)
	err = jsonParser.Decode(&packageJson)
	if err != nil {
		return errors.New("the package.json file is not properly formatted")
	}
	node.PackageJson.Scripts.Build = packageJson.Scripts.Build
	node.PackageJson.Scripts.Format = packageJson.Scripts.Format
	node.PackageJson.Scripts.Lint = packageJson.Scripts.Lint
	node.PackageJson.Scripts.Test = packageJson.Scripts.Test
	node.PackageJson.Scripts.TestUnit = packageJson.Scripts.TestUnit

	return nil
}

// defineLint enables the pipeline to execute the linter stage
func (node *Node) defineLint() {
	// Just execute the lint without install the packages
	if node.Manager == "None" {
		node.defaultLinter()
	} else {
		// Check for package.json
		if node.PackageJson.Scripts.Lint != "" {
			// run lint script
			node.Javascript.Language.Command.Linter = append(
				node.Javascript.Language.Command.Linter,
				node.normalizeManagerName(node.Manager) + " run lint",
			)
		} else {
			// There's no package, using stand alone config
			node.defaultLinter()
		}
	}
}

// defaultLinter is the default rule used in the linter stage
func (node *Node) defaultLinter() {
	node.Javascript.Language.Command.Linter = append(
		node.Javascript.Language.Command.Linter,
		"npx eslint --no-eslintrc --ext .js --ext .ts .",
	)
}

// defineFormat enables the pipeline to execute the format stage
func (node *Node) defineFormat() {
	if node.Manager == "None" {
		node.defaultFormat()
	} else {
		if node.PackageJson.Scripts.Lint != "" {
			// run lint script
			node.Javascript.Language.Command.Formater = append(
				node.Javascript.Language.Command.Formater,
				node.normalizeManagerName(node.Manager) + ` prettier --no-error-on-unmatched-pattern --check "**/*.js" "**/*.ts"`,
			)
		} else {
			// There's no package, using stand alone config
			node.defaultFormat()
		}
	}
}

// defaultFormat is the default rule used in the formatter stage
func (node *Node) defaultFormat() {
	node.Javascript.Language.Command.Formater = append(
		node.Javascript.Language.Command.Formater,
		`npx prettier --no-error-on-unmatched-pattern --check "**/*.js" "**/*.ts"`,
	)
}

// defineTest enables the pipeline to execute the test stage
func (node *Node) defineTest() {
	if node.Manager != "None" {
		if node.PackageJson.Scripts.Test != "" {
			node.Javascript.Language.Command.Test = append(
				node.Javascript.Language.Command.Test,
				node.normalizeManagerName(node.Manager) + " run test",
			)
		} else if node.PackageJson.Scripts.TestUnit != ""  {
			node.Javascript.Language.Command.Test = append(
				node.Javascript.Language.Command.Test,
				node.normalizeManagerName(node.Manager) + " run test:unit",
			)
		} else {
			node.defaultTest()
		}
	} else {
		node.defaultTest()
	}
}

// defaultTest is the default rule used in the test stage
func (node *Node) defaultTest() {
	node.Javascript.Language.Command.Test = append(
		node.Javascript.Language.Command.Test,
		"echo 'to run the test, you must configurate a test runner in the project'",
	)
}

// normalizeManagerName keep only the tool name used in the stage
func (node *Node) normalizeManagerName(name string) string {
	return strings.Fields(name)[0]
}

// fill sets the parameter to generate pointer
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

	generate.Command.Build = node.Javascript.Language.Command.Build
}