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
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/utils"
)

type LanguageType int

const (
	PYTHON LanguageType = iota
	NODE
	DENO
	GO
)

func (lt LanguageType) String() string {
	return [...]string{
		"Python",
		"Node - as Javascript Runtime",
		"Deno - as Javascript Runtime",
	}[lt]
}

type LanguageInterface interface {
	Init(generate *application.Generate)
	defineManager()
	defineFramework()
	defineVersion()
	defineLint()
	defineTest()
	defineFormat()
	fill(generate *application.Generate)
}

type Command struct {
	Build    []string
	Linter   []string
	Formater []string
	Test     []string
}

type Language struct {
	Command *Command
}

// Init Language definitions
func Init(generate *application.Generate) {
	var language LanguageInterface
	targetPlatform := ""
	platformPrompt := &survey.Select{
		Message: "Choose the project language:",
		Help:    "By the chosen language, you will be prompted with the properly configurations.",
		Options: []string{PYTHON.String(), NODE.String(), DENO.String()},
	}
	err := survey.AskOne(platformPrompt, &targetPlatform)
	if err != nil {
		utils.Error(err.Error())
	}

	switch targetPlatform {
	case PYTHON.String():
		language = NewPython()
		language.Init(generate)
	case NODE.String():
		language = NewJavascript(NODE.String())
		language.Init(generate)
	case DENO.String():
		language = NewJavascript(DENO.String())
		language.Init(generate)
	case GO.String():
		log.Println("Coming soon")
		os.Exit(0)
	default:
		utils.Error("language")
	}
}
