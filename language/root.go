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