package language

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/src/language/nodejs"
	"github.com/nogsantos/sarue/src/language/python"
	language "github.com/nogsantos/sarue/src/language/domain"
)

func Init() {
	targetPlatform := ""
	platformPrompt := &survey.Select{
		Message: "Choose the project language:",
		Options: language.Types(),
	}
	survey.AskOne(platformPrompt, &targetPlatform)

	switch targetPlatform {
		case language.PYTHON.String():
			python.Init()
			return
		case language.NODEJS.String():
			nodejs.Init()
			return
		default:
			return
	}
}