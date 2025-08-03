package internal

import (
	"bytes"
	"log/slog"
	"os"
)

func GetPrompt(promptFilename string) string {
	prompt, err := os.OpenFile(promptFilename, os.O_RDONLY, 0644)
	if err != nil {
		slog.Error("Error reading prompt", "err", err)
	}
	var buf bytes.Buffer
	buf.ReadFrom(prompt)
	return buf.String()
}
