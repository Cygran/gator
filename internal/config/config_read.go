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
	contents, err := os.Open(configFile)
	if err != nil {
		return Config{}, errors.New("error reading config file")
	}
	defer contents.Close()
	decoder := json.NewDecoder(contents)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
