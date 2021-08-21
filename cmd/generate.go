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
		generate := application.Generate{}
		generate.InitProcess()

		gitChecker()

		language.InitLanguage(&generate)
		platform.InitPlatform(&generate)

		generate.Create()
	},
}

func gitChecker() {
	// @todo: Remove-me
	// dur := time.Duration(rand.Intn(1000)) * time.Millisecond
	// time.Sleep(dur)

	git := conf.Git{}
	git.Init()
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
