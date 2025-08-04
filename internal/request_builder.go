package internal

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"os"
	"text/template"
)

type TemplateData struct {
	Message string
}

func BuildRequestJSON(messageToAI string) string {
	cfg := GetConfig()
	encodedMessage, err := json.Marshal(messageToAI)
	if err != nil {
		slog.Error("Error encoding request body", "error", err)
		os.Exit(1)
	}
	t, err := template.New("request").Parse(cfg.RequestTemplate)
	if err != nil {
		slog.Error("Error parse request template based on config", "error", err)
		os.Exit(1)
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, TemplateData{string(encodedMessage)}); err != nil {
		slog.Error("Error getting request data template", "error", err)
		os.Exit(1)
	}
	return buf.String()
}
