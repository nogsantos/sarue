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
	construct()
	Init(generate *application.Generate)
}

// Init Platform definitions
func Init(generate *application.Generate) {
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
		gl := NewGitLab()
		gl.Init(generate)
	case GITHUB.String():
		gb := NewGitHub()
		gb.Init(generate)
	default:
		utils.Error("platform")
	}
}
