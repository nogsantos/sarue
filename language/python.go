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
	LintCommands []string
	FormatCommands []string
	TestCommands []string
}

func (python *Python) construct() {
	python.versions = []string{"3.10", "3.9", "3.8", "3.7", "3.6", "3.5"}
	python.managers = []string{"pip", "pipenv", "poetry", "None"}
	python.frameworks = []string{"Django", "Pytests", "None"}
	python.stages = []string{"lint", "format", "test"}
	python.GithubActionsUser = "actions/setup-python@v2"
	python.GitLabBuildImage = ""
	python.LintCommands = []string{""}
	python.FormatCommands = []string{""}
	python.TestCommands = []string{""}
}

func (python *Python) Init(generate *application.Generate) {
	python.construct()
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
	python.TestCommands = append(python.TestCommands, "python -m pip install --upgrade pip")
	switch targetManager {
		case "pip":
			python.TestCommands = append(python.TestCommands, "python -m pip install -r requirements.txt")
		case "pipenv":
			python.TestCommands = append(python.TestCommands, "python -m pip install pipenv==2020.8.13")
			python.TestCommands = append(python.TestCommands, "pipenv install --system --deploy --python " + python.Version)
		case "poetry":
			python.TestCommands = append(python.TestCommands, "python -m pip install poetry")
			python.TestCommands = append(python.TestCommands, "poetry config virtualenvs.create false")
			python.TestCommands = append(python.TestCommands, "poetry install --no-root")
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
		Options: python.stages[:],
	}
	survey.AskOne(prompt, &targetStages)
	python.DefinedStages = targetStages
}

func (python *Python) defineLint() {
	python.LintCommands = append(python.LintCommands, "python -m pip install flake8")
	python.LintCommands = append(python.LintCommands, "flake8 . --ignore E203,E501,W503 --count --select=E9,F63,F7,F82 --show-source --exit-zero --max-complexity=10 --max-line-length=127 --statistics")
}

func (python *Python) defineFormat() {
	python.FormatCommands = append(python.FormatCommands, "python -m pip install black")
	python.FormatCommands = append(python.FormatCommands, "black . --check")
}

func (python *Python) defineTest() {
	switch python.Framework {
		case "Django":
			python.TestCommands = append(python.TestCommands, "python ./manage.py test --noinput --failfast -v 2")
		case "Pytests":
			python.TestCommands = append(python.TestCommands, "pytest -vv -s --log-level=INFO")
		default:
			python.TestCommands = append(python.TestCommands, "python -m pip pip install -r requirements.txt")
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
	generate.LintCommands = python.LintCommands
	generate.FormatCommands = python.FormatCommands
	generate.TestCommands = python.TestCommands
}