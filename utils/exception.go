package utils

import (
	"log"
	"os"
)

func Error(message string) {
	log.Fatalf("%v: A fatal error occour", message)
	os.Exit(0)
}