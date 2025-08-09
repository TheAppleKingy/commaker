package internal

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"regexp"

	"github.com/tidwall/gjson"
)

func GetCommitMessage(gitDiff string) string {
	cfg := GetConfig()
	msg := PROMPT + gitDiff
	jsonData := BuildRequestJSON(msg)
	request, err := http.NewRequest("POST", cfg.ProviderAPI, bytes.NewBufferString(jsonData))
	if err != nil {
		slog.Error("Cannot build request", "error", err)
		os.Exit(1)
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("COMMAKER_AI_API_KEY")))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		slog.Error("Cannot get response from AI model", "error", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error("Cannot read response body", "error", err)
		os.Exit(1)
	}
	re := regexp.MustCompile(`(?s)<think>.*?</think>`)
	answer := gjson.Get(string(bodyBytes), cfg.ResponsePath).String()
	cleaned := re.ReplaceAllString(answer, "")
	if cleaned == "" {
		slog.Error("Commit message was not marshalled", "model answer", string(bodyBytes))
		os.Exit(1)
	}
	return cleaned
}
