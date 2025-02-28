package config

import (
	"encoding/json"
	"errors"
	"os"
)

func Read() (Config, error) {
	configFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, errors.New("unable to get config file path")
	}
	contents, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, errors.New("error reading config file")
	}
	config := Config{}
	err = json.Unmarshal(contents, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
