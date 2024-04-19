package service

import (
	"fmt"
	"log"
	"os"

	"github.com/prtls/assets"
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

func readDirectory(directory string) ([]os.FileInfo, *os.File, error) {
	d, err := os.Open(directory)
	if err != nil {
		return nil, nil, err
	}

	files, err := d.Readdir(-1)
	if err != nil {
		return nil, nil, err
	}
	return files, d, nil
}

func closeDirectory(directory *os.File) {
	err := directory.Close()
	if err != nil {
		log.Fatalf("Error while closing directory: %s\n", err)
	}
}

func isHidden(name string) bool {
	return len(name) > 0 && name[0] == '.'
}

func formatFileDirectories(prefix string, file os.FileInfo) string {
	var format string
	if file.IsDir() {
		format = fmt.Sprintf("%s%s%s%s%s", prefix, assets.Bold, assets.Pink, file.Name(), assets.Reset)
	} else {
		format = fmt.Sprintf("%s%s", prefix, file.Name())
	}
	return format
}

func formatText(str string) string {
	return fmt.Sprintf("%s%s%s%s", assets.Bold, assets.LightBlue, str, assets.Reset)
}
