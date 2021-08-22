package language

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/utils"
)

type Node struct {
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
}

func (node *Node) construct() {
	node.platforms = []string{"Fronted", "Backend"}
	node.managers = []string{"Npm", "Yarn", "None"}
	node.frontFrameworks = []string{"Vue", "React", "None"}
	node.backendFrameworks = []string{"Nest", "Vuex", "Next"}
	node.nodeVersions = []string{"15.x", "14.x", "12.x", "10.x"}
	node.stages = []string{"lint", "format", "test"}
	node.GithubActionsUser = "actions/setup-node@v2"
	node.GitLabBuildImage = ""
}

func (node *Node) Init(generate *application.Generate) {
	node.construct()
	node.definePlatform()
	node.defineManager()
	// node.defineFrameworks()
	node.defineVersion()
	node.defineStages()

	node.fill(generate)
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

func (node *Node) defineVersion() {
	targetVersion := ""
	prompt := &survey.Select{
		Message: "What is the NodeJS version?",
		Options: node.nodeVersions[:],
	}
	err := survey.AskOne(prompt, &targetVersion)
	if err != nil {
		utils.Error(err.Error())
	}
	node.NodeVersion = targetVersion
}

func (node *Node) defineStages() {
	targetStages := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select the stages:",
		Options: node.stages[:],

	}
	survey.AskOne(prompt, &targetStages)
	node.DefinedStages = targetStages
}

func (node *Node) fill(generate *application.Generate) {
	generate.Language.Name = NODE.String()
	generate.ManagerName = node.Manager
	if node.FrontFramework != "" {
		generate.FrameworkName = node.FrontFramework
	} else {
		generate.FrameworkName = node.BackendFramework
	}
	generate.Language.Version = node.NodeVersion
	generate.DefinedBy = node.DefinedStages
	generate.Language.Extension = []string{"js", "ts"}
	generate.Language.GitHubCiBuilder = node.GithubActionsUser
	generate.Language.GitLabCiBuilder = node.GitLabBuildImage
}