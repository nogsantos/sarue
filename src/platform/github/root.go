package github

import (
	"fmt"
)


func Init() {

	initTarget := "github"

	// Platfoms["github"] = platform.Platfom{
	// 	Name: platform.GITHUB,
	// 	ConfigPath: "./.github/",
	// 	ConfigFile: "*",
	// 	Template: "",
	// }

	// pl, plError := Platfoms[initTarget].Validate()

	// if plError != nil {
	// 	fmt.Printf("Fudey")
	// 	os.Exit(0)
	// }

	fmt.Printf("initTarget %v", initTarget)
	// fmt.Printf("Create %v", pl.Name)
}
