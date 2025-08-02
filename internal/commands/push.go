package commands

import (
	"log/slog"
	"os"
	"os/exec"
)

func PushToRemote(localBranch, remoteBranch string) {
	cmd := exec.Command("git", "push", localBranch, remoteBranch)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		slog.Error("Error pushing: ", "error", err)
		os.Exit(1)
	}
}
