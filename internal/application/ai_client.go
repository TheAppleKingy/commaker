package application

import (
	"bytes"
	"commiter/internal"
	"commiter/internal/application/dto"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func GetCommitMessage(gitDiff string) string {
	prompt := internal.GetPrompt("prompt.template")
	data := dto.RequestData{
		Model: os.Getenv("MODEL"),
		Messages: []dto.Message{
			dto.Message{
				User:    "user",
				Content: fmt.Sprintf("%s\n%s", prompt, gitDiff),
			},
		},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Error("Encoding error", "error", err)
		os.Exit(1)
	}
	fmt.Print(string(jsonData))
	request, err := http.NewRequest("POST", os.Getenv("AI_MODEL_HOST"), bytes.NewBuffer(jsonData))
	if err != nil {
		slog.Error("Cannot build request", "error", err)
		os.Exit(1)
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPEN_AI_KEY")))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		slog.Error("Cannot get response from AI model", "error", err)
		os.Exit(1)
	}
	var respData map[string]interface{}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error("Cannot read response body", "error", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(bodyBytes, &respData); err != nil {
		slog.Error("Cannot parse response body", "error", err)
		os.Exit(1)
	}
	choices, ok := respData["choices"].([]interface{})
	if ok {
		choice, ok := choices[0].(map[string]interface{})
		if ok {
			message, ok := choice["message"].(map[string]interface{})
			if ok {
				commitMessage, ok := message["content"].(string)
				if ok {
					return commitMessage
				}
			}
		}
	}
	slog.Error("Unexpected response body", "invalid body", string(bodyBytes))
	os.Exit(1)
	return ""
}
