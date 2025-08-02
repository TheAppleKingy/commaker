package commands

import (
	"os"
)

// GetWorkingDir return directory where binary file was executed
func GetWorkingDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}
