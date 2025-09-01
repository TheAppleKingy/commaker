package commands

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

func GetCurrentBranch() string {
	cmd := exec.Command("git", "branch")
	var buf bytes.Buffer
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		slog.Error("Cannot get current git branch", "error", err)
		os.Exit(1)
	}
	var currentBranch string
	branches := strings.Split(buf.String(), "\n")
	for _, b := range branches {
		clear := strings.ReplaceAll(b, " ", "")
		if strings.HasPrefix(clear, "*") {
			currentBranch = string(clear[1:])
		}
	}
	return currentBranch
}
