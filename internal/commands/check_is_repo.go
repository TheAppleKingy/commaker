package commands

import (
	"log/slog"
	"os"
	"path/filepath"
)

// CheckIsLocalRepo checking that directory is git repository and set it as working directory and set it as working directory!
func CheckIsRepo() {
	matches, err := filepath.Glob(".git")
	if err != nil {
		slog.Error("Error finding .git in directory")
		os.Exit(1)
	}
	if len(matches) < 1 {
		slog.Error("Current directory is not a git repository")
		os.Exit(1)
	}
	CheckRemoteRepoExists()
}
