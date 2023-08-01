package fileutil

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFolderName string = ".cli-local-store"
const configFileName string = "local-store.json"

func fileExists(filePath string) bool {
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
	configFileDir := filepath.Dir(configFilePath)
	data := make(map[string]interface{})

	if fileExists(configFilePath) {
		// read the current content in data
		fmt.Println("Json file already exists")
	}

	jsonData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		fmt.Println("Error encoding json file: ", err)
		return
	}

	if err := os.MkdirAll(configFileDir, 0755); err != nil {
		fmt.Println("Error creating the folder path")
		return
	}

	if err := os.WriteFile(configFilePath, jsonData, 0644); err != nil {
		fmt.Println("Error writing json file", err)
		return
	}

	fmt.Println("Json file created")
}
