package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}
	
	return config, nil
}

func (c *Config)SetUser(userName string) (error) {
	c.CurrentUserName = userName

	if err := write(*c); err != nil {
		return err
	}
	return nil
}

func write(config Config) (error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, data, 0666); err != nil {
		return nil
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFileName), nil
}