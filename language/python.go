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
	Commands map[string][]string
}

func (python *Python) construct() {
	python.versions = []string{"3.10", "3.9", "3.8", "3.7", "3.6", "3.5"}
	python.managers = []string{"pip", "pipenv", "poetry", "None"}
	python.frameworks = []string{"Django", "Pytests", "None"}
	python.stages = []string{"lint", "format", "test[TDD]"}
	python.GithubActionsUser = "actions/setup-python@v2"
	python.GitLabBuildImage = ""
	// Its required initilize the commands map
	python.Commands = make(map[string][]string)
}

func (python *Python) Init(generate *application.Generate) {
	python.construct()
	python.defineVersion()
	python.defineManager()
	python.defineFramework()
	python.defineStages()
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
	python.Commands["upgrade"] = append(python.Commands["upgrade"], "python -m pip install --upgrade pip")
	switch targetManager {
		case "pip":
			python.Commands["install"] = append(python.Commands["install"], "python -m pip install -r requirements.txt")
		case "pipenv":
			python.Commands["install"] = append(python.Commands["install"], "python -m pip install pipenv==2020.8.13")
			python.Commands["install"] = append(python.Commands["install"], "pipenv install --system --deploy --python " + python.Version)
		case "poetry":
			python.Commands["install"] = append(python.Commands["install"], "python -m pip install poetry")
			python.Commands["install"] = append(python.Commands["install"], "poetry config virtualenvs.create false")
			python.Commands["install"] = append(python.Commands["install"], "poetry install --no-dev --no-root")
		default:
			python.Commands["install"] = append(python.Commands["install"], "python -m pip pip install -r requirements.txt")
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
	switch targetFramework {
		case "Django":
			python.Commands["test"] = append(python.Commands["test"], "manage.py test --noinput --failfast -v 2")
		case "Pytests":
			python.Commands["test"] = append(python.Commands["test"], "pytest -vv -s --log-level=INFO")
		default:
			python.Commands["test"] = append(python.Commands["test"], "python -m pip pip install -r requirements.txt")
	}
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
	generate.Commands = python.Commands
}