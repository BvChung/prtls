package fstraversal

import (
	"log"
	"os"
)

func CheckDefaultDirectory(directory *string) {
	if *directory == "" {

		var err error
		*directory, err = os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current working directory: %s\n", err)
		}
	}
}
