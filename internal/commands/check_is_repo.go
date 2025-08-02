package commands

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

// CheckIsLocalRepo checking that provided directory is git repository and set it as working directory
func CheckIsRepo(repoDir string) {
	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		slog.Error("Got not exsiting repository directory: " + repoDir)
		os.Exit(1)
	}
	matches, err := filepath.Glob(repoDir + "/.git")
	if err != nil {
		slog.Error("Error finding .git in " + repoDir)
		os.Exit(1)
	}
	fmt.Println(repoDir, matches)
	if len(matches) < 1 {
		slog.Error("Not found .git in " + repoDir + ". It is not repository")
		os.Exit(1)
	}
	os.Chdir(repoDir)
	CheckRemoteRepoExists()
}
