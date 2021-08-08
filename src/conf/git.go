package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// Git confirms if the repository has a valid git configuration
func Git() {
	content, readFileError := ioutil.ReadFile(".git/config")

	if readFileError != nil {
		initRepo()
	}

	hasRemoteEntry := strings.Contains(string(content), "remote \"origin\"")

	if hasRemoteEntry {
		fmt.Println("hasRemoteEntry")
	} else {
		fmt.Println("dont hasRemoteEntry")
	}
}

// init Initialize a empty repo by user confirmation
func initRepo() {
	initRepo := false
	prompt := &survey.Confirm{
		Message: "No git configuration file was found. Do you want to initialize it for you?",
		Default: true,
	}
	survey.AskOne(prompt, &initRepo)

	if !initRepo {
		fmt.Println("To create the configuration you must initialize your repo before.")
		os.Exit(0)
	}

	fmt.Println("Initializing the repo...")
	cmd := exec.Command("git", "init")
    stdout, commandError := cmd.Output()
	if commandError != nil {
		fmt.Printf("Git error %v", commandError)
		os.Exit(0)
	}
	fmt.Printf("%v", string(stdout))
}
