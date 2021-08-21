package platform

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/utils"
)

type GitLab struct {
	ConfigFile string
	ConfigPath string

}

func (gl *GitLab) Construct() {
	gl.ConfigFile = ".gitlab-ci.yml"
	gl.ConfigPath = "./"
}

func (gl *GitLab) Init(generate *application.Generate) {
	gl.Construct()
	gl.gitlabConfigExists(generate)
	gl.configureGitlaCi(generate)
}

func (gl *GitLab) gitlabConfigExists(generate *application.Generate) {
	path := fmt.Sprintf("%v/%v", utils.PathRoot(), gl.ConfigFile)
	_, readFileError := ioutil.ReadFile(path)
	if readFileError == nil {
		gl.confirmReplacement(generate)
	}
}

func (gl *GitLab) confirmReplacement(generate *application.Generate) {
	relplaceConfig := false
	prompt := &survey.Confirm{
		Message: "A Gitlab CI configuration was found, do you want to replace it?",
		Default: false,
	}
	err := survey.AskOne(prompt, &relplaceConfig)
	if err != nil {
		utils.Error(err.Error())
	}

	if !relplaceConfig {
		fmt.Printf("Finishing the process. We have nothing to do here.")
		os.Exit(0)
	} else {
		gl.configureGitlaCi(generate)
	}
}

func (gl *GitLab) configureGitlaCi(generate *application.Generate) {
	utils.WriteConfigFile(gl.ConfigPath, gl.ConfigFile, gl.GitlabTemplate(), generate)
}

func (gl *GitLab) GitlabTemplate() []byte {
	return []byte(`#===================================================================
# {{ .Tool }}
#
# {{ .Copyright }}
#
{{ if .Legal.Text }}{{ .Legal.Text }}{{ end }}
#===================================================================
stages:
  - test
  - build
  - deploy

`)
}