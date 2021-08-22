package language

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/utils"
)

type LanguageType int
const (
	PYTHON LanguageType = iota
	NODE
)
func (lt LanguageType) String() string {
	return [...]string{"Python", "Node"}[lt]
}

type LanguageInterface interface {
	Init(generate *application.Generate)
	construct()
	defineManager()
	defineFramework()
	defineVersion()
	fill(generate *application.Generate)
}

func InitLanguage(generate *application.Generate) {
	targetPlatform := ""
	platformPrompt := &survey.Select{
		Message: "Choose the project language:",
		Options: []string{PYTHON.String(), NODE.String()},
	}
	err := survey.AskOne(platformPrompt, &targetPlatform)
	if err != nil {
		utils.Error(err.Error())
	}

	switch targetPlatform {
		case PYTHON.String():
			py := Python{}
			py.Init(generate)
		case NODE.String():
			nd := Node{}
			nd.Init(generate)
		default:
			utils.Error("Language")
	}
}