package utils

import (
	"fmt"
	"os"
)

// PathRoot Gets the root of started process on cli
func PathRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error to found path %v", err)
		os.Exit(0)
	}
	return wd
}