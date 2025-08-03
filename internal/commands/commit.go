package commands

import (
	"log/slog"
	"os"
	"os/exec"
)

func Commit(commitMessage string) {
	cmd := exec.Command("git", "commit", "-m", commitMessage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		slog.Error("Error commiting changes", "error", err)
		os.Exit(1)
	}
}
