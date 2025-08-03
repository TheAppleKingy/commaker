package commands

import (
	"log/slog"
	"os"
	"os/exec"
)

func PushToRemote(branch string) {
	if branch == "" {
		branch = DetectCurrentBranch()
	}
	args := []string{"push", "origin", branch}
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		slog.Error("Error pushing: ", "error", err)
		os.Exit(1)
	}
}
