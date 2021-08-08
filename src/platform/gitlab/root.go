package gitlab

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/src/utils"
)

const configFile string = ".gitlab-ci.yml"

func Init() {
	gitlabConfigExists()
	configureGitlaCi()
}

func gitlabConfigExists() {
	path := fmt.Sprintf("%v/%v", utils.PathRoot(), configFile)
	_, readFileError := ioutil.ReadFile(path)
	if readFileError == nil {
		confirmReplacement()
	}
}

func confirmReplacement() {
	relplaceConfig := false
	prompt := &survey.Confirm{
		Message: "A Gitlab CI configuration was found, do you want to replace it?",
		Default: false,
	}
	survey.AskOne(prompt, &relplaceConfig)

	if !relplaceConfig {
		fmt.Printf("Finishing the process. We have nothing to do here.")
		os.Exit(0)
	} else {
		configureGitlaCi()
	}
}

func configureGitlaCi() {
	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile(configFile, message, 0644)
	if err != nil {
		fmt.Println("Will configure gitlab")
		os.Exit(0)
	}
}
