package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
)

const configFolderName string = ".cli-local-store"
const configFileName string = "local-store.json"

func existFile(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	configFilePath := filepath.Join(homeDir, configFolderName, configFileName)

	return configFilePath, nil
}

func Init() {
	configFilePath, _ := getConfigFilePath()

	fmt.Println(existFile(configFilePath))
}
