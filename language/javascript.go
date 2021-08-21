package language

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/utils"
)

type Javascript struct {
	platforms []string
	managers []string
	frontFrameworks []string
	backendFrameworks []string
	nodeVersions []string
	stages []string
	Platform string
	Manager string
	FrontFramework string
	BackendFramework string
	NodeVersion string
	DefinedStages []string
	GithubActionsUser string
	GitLabBuildImage string
	Runtimer string
}

func (javascript *Javascript) construct() {
	javascript.platforms = []string{"Fronted", "Backend"}
	javascript.managers = []string{"Npm", "Yarn", "None"}
	javascript.frontFrameworks = []string{"Vue", "React", "None"}
	javascript.backendFrameworks = []string{"Nest", "Vuex", "Next"}
	javascript.nodeVersions = []string{"v16", "v14", "v12"}
	javascript.stages = []string{"lint", "format", "test"}
	javascript.GithubActionsUser = "actions/javascript-action@v1"
	javascript.GitLabBuildImage = ""
	javascript.Runtimer = "node"
}

func (javascript *Javascript) Init(generate *application.Generate) {
	javascript.construct()
	javascript.definePlatform()
	javascript.defineManager()
	// javascript.defineFrameworks()
	javascript.defineVersion()
	javascript.defineStages()

	javascript.fill(generate)
}

func (javascript *Javascript) definePlatform() {
	targetPlatform := ""
	platformPrompt := &survey.Select{
		Message: "Will be used for front or backend:",
		Options: javascript.platforms[:],
	}
	err := survey.AskOne(platformPrompt, &targetPlatform)
	if err != nil {
		utils.Error(err.Error())
	}
	javascript.Platform = targetPlatform

	switch targetPlatform {
		case javascript.platforms[0]:
			javascript.frontendHandler()
		case javascript.platforms[1]:
			javascript.backendHandler()
	default:
		utils.Error("Javascript")
	}
}

func (javascript *Javascript) frontendHandler() {
	targetFramework := ""
	prompt := &survey.Select{
		Message: "What is the framework:",
		Options: javascript.frontFrameworks[:],
	}
	err := survey.AskOne(prompt, &targetFramework)
	if err != nil {
		utils.Error(err.Error())
	}
	javascript.BackendFramework = ""
	javascript.FrontFramework = targetFramework
}

func (javascript *Javascript) backendHandler() {
	targetFramework := ""
	prompt := &survey.Select{
		Message: "What is the framework:",
		Options: javascript.backendFrameworks[:],
	}
	err := survey.AskOne(prompt, &targetFramework)
	if err != nil {
		utils.Error(err.Error())
	}
	javascript.FrontFramework = ""
	javascript.BackendFramework = targetFramework
}

func (javascript *Javascript) defineManager() {
	targetManager := ""
	prompt := &survey.Select{
		Message: "Using a manager?",
		Options: javascript.managers[:],
	}
	err := survey.AskOne(prompt, &targetManager)
	if err != nil {
		utils.Error(err.Error())
	}
	javascript.Manager = targetManager
}

func (javascript *Javascript) defineVersion() {
	targetVersion := ""
	prompt := &survey.Select{
		Message: "What is the python version?",
		Options: javascript.nodeVersions[:],
	}
	err := survey.AskOne(prompt, &targetVersion)
	if err != nil {
		utils.Error(err.Error())
	}
	javascript.NodeVersion = targetVersion
}

func (javascript *Javascript) defineStages() {
	targetStages := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select the stages:",
		Options: javascript.stages[:],

	}
	survey.AskOne(prompt, &targetStages)
	javascript.DefinedStages = targetStages
}

func (javascript *Javascript) fill(generate *application.Generate) {
	generate.Language.Name = JAVASCRIPT.String()
	generate.ManagerName = javascript.Manager
	if javascript.FrontFramework != "" {
		generate.FrameworkName = javascript.FrontFramework
	} else {
		generate.FrameworkName = javascript.BackendFramework
	}
	generate.Language.Version = javascript.NodeVersion
	generate.DefinedBy = javascript.DefinedStages
	generate.Language.Extension = []string{"js", "ts"}
	generate.Language.GitHubCiBuilder = javascript.GithubActionsUser
	generate.Language.GitLabCiBuilder = javascript.GitLabBuildImage
}