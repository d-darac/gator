package config

import (
	"encoding/json"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(cfg)
}

func write(cfg *Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	buf, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	if err = os.WriteFile(path, buf, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, configFileName), nil
}
