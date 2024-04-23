package fstraversal

import (
	"fmt"
	"os"
	"regexp"
)

func cleanTree(treeStructure string) string {
	return removeANSI(treeStructure)
}

func removeANSI(input string) string {
	reg := regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")
	return reg.ReplaceAllString(input, "")
}

func writeToFile(path string, data []byte) error {
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("unable to write to file: %w", err)
	}
	return nil
}
