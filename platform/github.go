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
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/utils"
)

type Github struct {
	Lang string
	ConfigDir string
	ConfigFile string
}

func (gb *Github) construct() {
	gb.ConfigDir = ".github/workflows/"
}

func (gb *Github) Init(generate *application.Generate) {
	gb.construct()
	gb.githubConfigExists()
	gb.ParseFile(generate)
}

func (gb *Github) githubConfigExists() {
	_, err := ioutil.ReadDir(gb.ConfigDir)
	if err != nil {
		gb.createConfDir()
	}

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }
}

func (gb *Github) createConfDir() {
	err := os.MkdirAll(gb.ConfigDir, 0755)
    if err != nil {
        log.Fatal(err)
    }
}

func (gb *Github) ParseFile(generate *application.Generate) {
	gb.ConfigFile = strings.ToLower(generate.Language.Name)
	gb.configureFile(generate)
}

func (gb *Github) configureFile(generate *application.Generate) {
	if len(generate.DefinedBy) > 0 {
		for _, conf := range generate.DefinedBy {
			tartargetConf := gb.GithubLintTemplate()
			if conf == "format" {
				tartargetConf = gb.GithubFormatTemplate()
			} else if conf == "test" {
				tartargetConf = gb.GithubTestTemplate()
			}
			utils.WriteConfigFile(gb.ConfigDir, gb.ConfigFile+"."+conf + ".yaml", tartargetConf, generate)
		}
	}
}
