package platform

import (
	"github.com/AlecAivazis/survey/v2"
	platform "github.com/nogsantos/sarue/src/platform/domain"
	"github.com/nogsantos/sarue/src/platform/github"
	"github.com/nogsantos/sarue/src/platform/gitlab"
)



func Init() platform.Platfom {
	targetPlatform := ""
	platformPrompt := &survey.Select{
		Message: "Choose the platform:",
		Options: platform.Types(),
	}
	survey.AskOne(platformPrompt, &targetPlatform)

	switch targetPlatform {
		case platform.GITLAB.String():
			pl := platform.Platfom{
				Name: platform.GITLAB,
			}

			gitlab.Init()
			return pl
		case platform.GITHUB.String():
			pl := platform.Platfom{
				Name: platform.GITLAB,
			}
			github.Init()
			return pl
		default:
			pl := platform.Platfom{}
			return pl
	}
}

// func (pl Platfom) Validate() (Platfom, error) {

// 	path := pl.ConfigPath
// 	file := pl.ConfigFile

// 	// files, dirEerror := ioutil.ReadDir(path)
// 	dirEerror := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println(path, info.Size())
// 		return nil
// 	})

// 	if dirEerror != nil {
// 		fmt.Printf("Path not found")
// 		os.Exit(0)
// 	}

// 	// for _, file := range files {
// 	// 	fmt.Println(file.Name())
// 	// }
// 	os.Exit(0)

// 	configFile, readFileError := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, file))

// 	// cobra.CheckErr(readFileError.Error())

// 	if readFileError != nil {
// 		fmt.Printf("Did not found the file, create it?")
// 		// tempMaker()
// 	} else {
// 		fmt.Printf("%v", string(configFile))

// 		relplaceConfig := false
// 		prompt := &survey.Confirm{
// 			Message: "A configuration file was found, do you wanna replace it?",
// 			Default: false,
// 		}
// 		survey.AskOne(prompt, &relplaceConfig)

// 		if !relplaceConfig {
// 			fmt.Printf("Finish the process.")
// 			os.Exit(0)
// 		} else {
// 			fmt.Printf("Replace it?")
// 		}
// 	}

// 	return pl, nil
// }

// func findConfigurationFile(source string) (bool, error) {
// 	matched, err := regexp.Match(`yml|yaml`, []byte(source))
// 	fmt.Println(matched, err)
// 	return true, nil
// }