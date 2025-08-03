package commands

import (
	"os"
	"os/exec"
)

func SwitchToBranch(branchName string) {
	cmd := exec.Command("git", "switch", branchName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
