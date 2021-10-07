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

type Python struct {
	Language *Language
	managers []string
	frameworks []string
	versions []string
	stages []string
	Manager string
	Framework string
	Version string
	DefinedStages []string
	GithubActionsUser string
	GitLabBuildImage string
}

// NewPython create a new Python instance
func NewPython() *Python {
	return &Python{
		versions: []string{"3.10", "3.9", "3.8", "3.7", "3.6", "3.5"},
		managers: []string{"pip", "pipenv", "poetry", "None"},
		frameworks: []string{"Django", "Pytests", "None"},
		stages: []string{"lint", "format", "test"},
		GithubActionsUser: "actions/setup-python@v2",
		GitLabBuildImage: "python:-alpine",
		Language: &Language{
			Command: &Command{
				Linter: []string{},
				Formater: []string{},
				Test: []string{},
			},
		},
	}
}

// Init Python configuration process
func (python *Python) Init(generate *application.Generate) {
	python.defineStages()
	python.defineVersion()
	for _, stage := range python.DefinedStages {
		if stage == "test" {
			python.defineManager()
			python.defineFramework()
			python.defineTest()
		}
		if stage == "lint" {
			python.defineLint()
		}
		if stage == "format" {
			python.defineFormat()
		}
	}
	python.fill(generate)
}

func (python *Python) defineVersion() {
	targetVersion := ""
	prompt := &survey.Select{
		Message: "What is the python version?",
		Options: python.versions[:],
	}
	err := survey.AskOne(prompt, &targetVersion)
	if err != nil {
		utils.Error(err.Error())
	}
	python.Version = targetVersion
}

func (python *Python) defineManager() {
	targetManager := ""
	prompt := &survey.Select{
		Message: "Using a manager?",
		Options: python.managers[:],
	}
	err := survey.AskOne(prompt, &targetManager)
	if err != nil {
		utils.Error(err.Error())
	}
	python.Manager = targetManager
	python.Language.Command.Test = append(python.Language.Command.Test, "python -m pip install --upgrade pip")
	switch targetManager {
		case "pip":
			python.Language.Command.Test = append(python.Language.Command.Test, "python -m pip install -r requirements.txt")
		case "pipenv":
			python.Language.Command.Test = append(python.Language.Command.Test, "python -m pip install pipenv==2020.8.13")
			python.Language.Command.Test = append(python.Language.Command.Test, "pipenv install --system --deploy --python " + python.Version)
		case "poetry":
			python.Language.Command.Test = append(python.Language.Command.Test, "python -m pip install poetry")
			python.Language.Command.Test = append(python.Language.Command.Test, "poetry config virtualenvs.create false")
			python.Language.Command.Test = append(python.Language.Command.Test, "poetry install --no-root")
	}
}

func (python *Python) defineFramework() {
	targetFramework := ""
	prompt := &survey.Select{
		Message: "Using a framework?",
		Options: python.frameworks[:],
	}
	err := survey.AskOne(prompt, &targetFramework)
	if err != nil {
		utils.Error(err.Error())
	}
	python.Framework = targetFramework
}

func (python *Python) defineStages() {
	targetStages := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select the stages:",
		Help: "Stages are the steps that the pipeline will cover.",
		Options: python.stages[:],
	}
	survey.AskOne(prompt, &targetStages)
	python.DefinedStages = targetStages
}

func (python *Python) defineLint() {
	python.Language.Command.Linter = append(python.Language.Command.Linter, "python -m pip install flake8")
	python.Language.Command.Linter = append(python.Language.Command.Linter, "flake8 . --ignore E203,E501,W503 --count --select=E9,F63,F7,F82 --show-source --exit-zero --max-complexity=10 --max-line-length=127 --statistics")
}

func (python *Python) defineFormat() {
	python.Language.Command.Formater = append(python.Language.Command.Formater, "python -m pip install black")
	python.Language.Command.Formater = append(python.Language.Command.Formater, "black . --check")
}

func (python *Python) defineTest() {
	switch python.Framework {
		case "Django":
		python.Language.Command.Test = append(python.Language.Command.Test, "python ./manage.py test --noinput --failfast -v 2")
		case "Pytests":
			python.Language.Command.Test = append(python.Language.Command.Test, "pytest -vv -s --log-level=INFO")
		default:
			python.Language.Command.Test = append(python.Language.Command.Test, "python -m pip pip install -r requirements.txt")
	}
}

func (python *Python) fill(generate *application.Generate) {
	generate.Language.Name = PYTHON.String()
	generate.ManagerName = python.Manager
	generate.FrameworkName = python.Framework
	generate.Language.Version = python.Version
	generate.DefinedBy = python.DefinedStages
	generate.Language.Extension = []string{"py"}
	generate.Language.GitHubCiBuilder = python.GithubActionsUser
	generate.Language.GitLabCiBuilder = python.GitLabBuildImage
	generate.Command.Linter = python.Language.Command.Linter
	generate.Command.Formatter = python.Language.Command.Formater
	generate.Command.Test =  python.Language.Command.Test
}