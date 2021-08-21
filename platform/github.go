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

func (gb *Github) Construct() {
	gb.ConfigDir = ".github/workflows/"
}

func (gb *Github) Init(generate *application.Generate) {
	gb.Construct()
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

func (gb *Github) GithubLintTemplate() []byte {
	return []byte(`#===================================================================
# {{ .Tool }}
#
# {{ .Copyright }}
#
{{ if .Legal.Text }}{{ .Legal.Text }}{{ end }}
#===================================================================
name: Lint {{ .Language.Name }}
on:
  pull_request:
    paths:{{ range .Language.Extension }}
      - '**.{{ . }}'{{ end }}

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Set up {{ .Language.Name }}
        uses: {{ .Language.GitHubCiBuilder }}
        with:
          python-version: {{ .Language.Version }}
      - run: python -m pip install flake8
      - run: flake8

`)
}

func (gb *Github) GithubTestTemplate() []byte {
	return []byte(`#===================================================================
# {{ .Tool }}
#
# {{ .Copyright }}
#
{{ if .Legal.Text }}{{ .Legal.Text }}{{ end }}
#===================================================================
name: Test {{ .Language.Name }}
on:
  pull_request:
    paths:{{ range .Language.Extension }}
      - '**.{{ . }}'{{ end }}

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Set up {{ .Language.Name }}
        uses: {{ .Language.GitHubCiBuilder }}
        with:
          python-version: {{ .Language.Version }}
      - run: python -m pip install flake8
      - run: flake8

`)
}

func (gb *Github) GithubFormatTemplate() []byte {
	return []byte(`#===================================================================
# {{ .Tool }}
#
# {{ .Copyright }}
#
{{ if .Legal.Text }}{{ .Legal.Text }}{{ end }}
#===================================================================
name: Format {{ .Language.Name }}
on:
  pull_request:
    paths:{{ range .Language.Extension }}
      - '**.{{ . }}'{{ end }}

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Set up {{ .Language.Name }}
        uses: {{ .Language.GitHubCiBuilder }}
        with:
          python-version: {{ .Language.Version }}
      - run: python -m pip install flake8
      - run: flake8

`)
}
