package platform

import (
	"fmt"
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
	files, err := ioutil.ReadDir(gb.ConfigDir)
	if err != nil {
		gb.createConfDir()
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func (gb *Github) createConfDir() {
	err := os.MkdirAll(gb.ConfigDir, 0755)
    if err != nil {
        log.Fatal(err)
    }
}

func (gb *Github) ParseFile(generate *application.Generate) {
	gb.ConfigFile = strings.ToLower(generate.Language.Name)
	fmt.Printf("TEST1 %v \n",gb.ConfigFile)
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
			fmt.Printf("TEST %v \n",gb.ConfigFile)
			utils.WriteConfigFile(gb.ConfigDir, gb.ConfigFile+"."+conf + ".yaml", tartargetConf, generate)
		}
	}
}
