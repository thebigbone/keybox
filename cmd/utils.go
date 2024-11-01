package main

import (
	"os"
)

func getPath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	keyStorePath := homedir + "/.keybox"

	return keyStorePath, nil
}
