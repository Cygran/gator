package config

import (
	"errors"
	"os"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("unable to get home directory")
	}
	return homeDir + "/" + configFileName, nil
}
