package cmd

import (
	"fmt"

	"github.com/nogsantos/sarue/src/conf"
	"github.com/nogsantos/sarue/src/language"
	"github.com/nogsantos/sarue/src/platform"
	"github.com/spf13/cobra"
)

var generate conf.Generate

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Start a CI configuration",
	Long: `Init a configuration base file to your project.`,
	Run: func(cmd *cobra.Command, args []string) {
		generate = conf.Generate{
			InitAt: "Now",
		}

		conf.Git()

		pl := platform.Init()
		language.Init()
		generateInit()

		generate.FinishAt = "some"
		generate.Platfom = pl
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateInit() {
	fmt.Printf("\nWill generate the configuration: %v %v", generate.Platfom.Name, generate.FinishAt)
}
