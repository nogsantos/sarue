package language

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/utils"
)

type Python struct {
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

func (python *Python) construct() {
	python.managers = []string{"pip", "pyenv", "poetry", "None"}
	python.frameworks = []string{"Django", "Pytests", "None"}
	python.versions = []string{"3.10", "3.9", "3.8", "3.7", "3.6", "3.5"}
	python.stages = []string{"lint", "format", "test"}
	python.GithubActionsUser = "actions/setup-python@v2"
	python.GitLabBuildImage = ""
}

func (python *Python) Init(generate *application.Generate) {
	python.construct()
	python.defineManager()
	python.defineFramework()
	python.defineVersion()
	python.defineStages()
	python.fill(generate)
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

func (python *Python) defineStages() {
	targetStages := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select the stages:",
		Options: python.stages[:],
	}
	survey.AskOne(prompt, &targetStages)
	python.DefinedStages = targetStages
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
}