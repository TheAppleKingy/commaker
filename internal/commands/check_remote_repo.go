package commands

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

// CheckRemoteRepoExists check that local repository is related with remote. Function should be called in local repo directory
func CheckRemoteRepoExists() {
	cmd := exec.Command("git", "remote", "-v")
	var remoteInfo bytes.Buffer
	var errorMessage bytes.Buffer
	cmd.Stdout = &remoteInfo
	cmd.Stderr = &errorMessage
	if err := cmd.Run(); err != nil {
		slog.Error("Error checking remote repo", "error", errorMessage.String())
		os.Exit(1)
	}
	if remoteInfo.String() == "" {
		slog.Error("Local repository not related with any remote repository")
		os.Exit(1)
	}
	remotes := strings.Split(remoteInfo.String(), "\n")
	for _, remote := range remotes {
		if strings.HasPrefix(remote, "origin") && strings.HasSuffix(remote, "(push)") {
			return
		}
	}
	slog.Error("Incorrect info about remote repo", "unexpected info", remoteInfo.String())
	os.Exit(1)
}
