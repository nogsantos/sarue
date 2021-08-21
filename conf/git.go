package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nogsantos/sarue/utils"
)

type Git struct {}

// Git confirms if the repository has a valid git configuration
func (git *Git) Init() {
	_, readFileError := ioutil.ReadFile(".git/config")

	if readFileError != nil {
		git.initRepo()
	}

	// hasRemoteEntry := strings.Contains(string(content), "remote \"origin\"")

	// if hasRemoteEntry {
	// 	fmt.Println("hasRemoteEntry")
	// } else {
	// 	fmt.Println("dont hasRemoteEntry")
	// }
}

// init Initialize a empty repo by user confirmation
func (git *Git) initRepo() {
	initRepo := false
	prompt := &survey.Confirm{
		Message: "No git configuration file was found. Do you want us to initialize it for you?",
		Default: true,
	}
	err := survey.AskOne(prompt, &initRepo)
	if err != nil {
		utils.Error(err.Error())
	}

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
