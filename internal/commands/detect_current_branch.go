package commands

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

func DetectCurrentBranch() string {
	cmd := exec.Command("git", "branch")
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Run()
	var currentBranch string
	branches := strings.Split(buf.String(), "\n")
	for _, branchName := range branches {
		if strings.HasPrefix(branchName, "*") {
			currentBranch = strings.Split(branchName, " ")[1]
		}
	}
	if currentBranch == "" {
		slog.Error("Unexpected format of info about branches", "invalid data", strings.Join(branches, "\n"))
		os.Exit(1)
	}
	return currentBranch
}
