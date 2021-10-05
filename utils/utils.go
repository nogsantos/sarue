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

package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	cmdFile, err := os.Create(fmt.Sprintf("%s%s", path, fileName))
	if err != nil {
		Error("Create file " + err.Error())
	}
	defer cmdFile.Close()

	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
	}


	githubTemplate := template.Must(template.New(data.Language.Name).Funcs(funcMap).Parse(string(local_template)))

	err = githubTemplate.Execute(cmdFile, data)
	if err != nil {
		Error(err.Error())
	}
}
