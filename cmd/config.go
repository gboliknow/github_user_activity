package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	GitHubToken  string `json:"github_token"`
	DefaultUser  string `json:"default_user"`
	CacheTTL     int    `json:"cache_ttl"`
	OutputFormat string `json:"output_format"`
}

var config Config

func loadConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(home, ".github_user_activity.json")
	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Use defaults
		}
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&config)
}

func saveConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(home, ".github_user_activity.json")
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(config)
}
