package commands

import (
	"log/slog"
	"os"
	"os/exec"
)

func PushToRemote() {
	cmd := exec.Command("git", "push")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		slog.Error("Error pushing: ", "error", err)
		os.Exit(1)
	}
}
