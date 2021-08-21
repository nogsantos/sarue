package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/nogsantos/sarue/application"
)

// PathRoot Gets the root of started process on cli
func PathRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		errorMessage := "Error to found path " + err.Error()
		Error(errorMessage)
	}
	return wd
}

func CreateConfigFile(path, fileName, content string) []byte {
	bcontent := []byte(content)
	err := ioutil.WriteFile(path+fileName, bcontent, 0644)
	if err != nil {
		errorMessage := "Create file " + err.Error()
		Error(errorMessage)
	}

	file, _ := ioutil.ReadFile(path+fileName)
	return file
}

func WriteConfigFile(path, fileName string, local_template []byte, data *application.Generate) {
	fmt.Print(fileName)
	cmdFile, err := os.Create(fmt.Sprintf("%s%s", path, fileName))
	if err != nil {
		Error("Create file " + err.Error())
	}
	defer cmdFile.Close()

	githubTemplate := template.Must(template.New(data.Language.Name).Parse(string(local_template)))

	err = githubTemplate.Execute(cmdFile, data)
	if err != nil {
		Error(err.Error())
	}
}
