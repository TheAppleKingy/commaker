package commands

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
)

func GetGitDiff() string {
	cmd := exec.Command("git", "diff", "--staged")
	var changes bytes.Buffer
	cmd.Stdout = &changes
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		slog.Error("Cannot get code changes", "error", err)
		os.Exit(1)
	}
	if changes.String() == "" {
		slog.Error("No changes staged")
		os.Exit(1)
	}
	return changes.String()
}
