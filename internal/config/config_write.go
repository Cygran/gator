package config

import (
	"encoding/json"
	"os"
)

func write(cfg *Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
