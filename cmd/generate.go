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

package cmd

import (
	"github.com/nogsantos/sarue/application"
	"github.com/nogsantos/sarue/conf"
	"github.com/nogsantos/sarue/language"
	"github.com/nogsantos/sarue/platform"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Long: `Init a configuration base file to your project.`,
	Short: "Start a CI configuration",
	Aliases: []string{"g"},
	Example: `
  Long form: sarue generate
  Short form: sarue g`,
	Run: func(cmd *cobra.Command, args []string) {
		generate := *application.NewGenerate()
		// Validate if is a git repository
		gitChecker()
		// Setting up the languages
		language.Init(&generate)
		// Setting up the platform
		platform.Init(&generate)
		// Finishing the process
		generate.Finish()
	},
}

func gitChecker() {
	git := conf.Git{}
	git.Init()
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
