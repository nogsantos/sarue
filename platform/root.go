package platform

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/utils"
)

type PlatfomType int
const (
	GITHUB PlatfomType = iota
	GITLAB
)

func (p PlatfomType) String() string {
	return [...]string{"Github", "Gitlab"}[p]
}

type Platfom interface {
	Construct()
	Init(generate *application.Generate)
}

func InitPlatform(generate *application.Generate) {
	targetPlatform := ""
	platformPrompt := &survey.Select{
		Message: "Choose the platform:",
		Options: []string{GITHUB.String(), GITLAB.String()},
	}
	err := survey.AskOne(platformPrompt, &targetPlatform)
	if err != nil {
		utils.Error(err.Error())
	}

	switch targetPlatform {
		case GITLAB.String():
			gl := GitLab{}
			gl.Init(generate)
		case GITHUB.String():
			gb := Github{}
			gb.Init(generate)
		default:
			utils.Error("Platform")
	}
}
