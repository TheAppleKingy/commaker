package internal

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ProviderAPI     string `yaml:"provider_api"`
	RequestTemplate string `yaml:"request_template"`
	ResponsePath    string `yaml:"response_path"`
}

func GetConfig() Config {
	confPath := os.Getenv("COMMAKER_CONFIG_PATH")
	if confPath == "" {
		slog.Error("Path to config was not set!")
		os.Exit(1)
	}
	file, err := os.ReadFile(confPath)
	if err != nil {
		slog.Error("Error reading template file", "error", err)
		os.Exit(1)
	}
	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		slog.Error("Error parse config", "error", err)
		os.Exit(1)
	}
	return config
}
