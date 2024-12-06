package helper

import (
	"os"
	"path/filepath"
)

func initializeUserConfig() error {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	configDir := filepath.Join(homeDir, ".cli-todo", "config.json")

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(filepath.Dir(configDir), 0755)

		if err != nil {
			return err
		}
	}

	return nil

}
