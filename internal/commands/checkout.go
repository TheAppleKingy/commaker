package commands

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
)

func SwitchToBranch(branchName string) {
	cmd := exec.Command("git", "switch", branchName)
	var buf bytes.Buffer
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		slog.Error("Error sdwitching to provided branch", "error", buf.String())
		os.Exit(1)
	}
}
