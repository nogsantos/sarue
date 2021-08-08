package python

import (
	"fmt"

	language "github.com/nogsantos/sarue/src/language/domain"
)

type PythonManager int
func (p PythonManager) String() string {
	return [...]string{"pip", "pyenv", "poetry"}[p]
}

type PythonFramework int
func (f PythonFramework) String() string {
	return [...]string{"Django", "Pytests"}[f]
}

// type Python struct {
// 	Manager packageManager
// 	Framework Framework
// }

func Init() {
	fmt.Println("Python initialized")
	pythonPackages := &language.PackageManager{
		Packages: []string{"pip", "pyenv", "poetry"},
	}
	fmt.Printf("Packages %v", *pythonPackages)
	pythonFramework := &language.Framework{
		Frameworks: []string{"Django", "Pytests"},
	}
	fmt.Printf("Frameworks %v", *pythonFramework)
	// python := &Language{
	// 	PackageManager
	// }
}