package localstoreutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const folderName string = ".cli-local-store"
const fileName string = "local-store.json"

func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func getFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	configFilePath := filepath.Join(homeDir, folderName, fileName)

	return configFilePath, nil
}

func createFileIfNotExist() {
	filePath, _ := getFilePath()
	fileDir := filepath.Dir(filePath)
	data := make(map[string]interface{})

	if fileExists(filePath) {
		return
	}

	jsonData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return
	}

	if err := os.MkdirAll(fileDir, 0755); err != nil {
		return
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return
	}
}

func StoreKey(key string, value string, force bool) (string, error) {
	createFileIfNotExist()

	filePath, _ := getFilePath()
	data := make(map[string]interface{})

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", errors.New("error reading the existing data")
	}

	if err := json.Unmarshal(fileContent, &data); err != nil {
		return "", errors.New("error parsing the existing dataa")
	}

	data[key] = value

	newData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return "", errors.New("error reading the existing data")
	}

	if err := os.WriteFile(filePath, newData, 0644); err != nil {
		return "", errors.New("error writing the new data")
	}

	return fmt.Sprintf("The key '%s' was store successfuly", key), nil
}
