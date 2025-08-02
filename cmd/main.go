package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		slog.Error("Error command executing")
		return
	}
	diff := out.String()
	fmt.Print(diff)
}
