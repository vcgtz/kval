package localstoreutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const folderName string = ".cli-kval"
const fileName string = "kval.json"

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

func getContent() (map[string]interface{}, error) {
	filePath, _ := getFilePath()
	data := make(map[string]interface{})

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("error reading the existing data")
	}

	if err := json.Unmarshal(fileContent, &data); err != nil {
		return nil, errors.New("error parsing the existing data")
	}

	return data, nil
}

func saveContent(data map[string]interface{}) error {
	filePath, _ := getFilePath()
	newData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return errors.New("error reading the existing data")
	}

	if err := os.WriteFile(filePath, newData, 0644); err != nil {
		return errors.New("error writing the new data")
	}

	return nil
}

func existKey(data map[string]interface{}, key string) bool {
	_, ok := data[key]

	return ok
}

func StoreValue(key string, value string, force bool) (string, error) {
	createFileIfNotExist()

	data, err := getContent()
	if err != nil {
		return "", err
	}

	if existKey(data, key) && !force {
		return "duplicated", nil
	}

	data[key] = value

	if err := saveContent(data); err != nil {
		return "", err
	}

	return "success", nil
}

func GetValue(key string) (string, error) {
	data, err := getContent()
	if err != nil {
		return "", err
	}

	if existKey(data, key) {
		return fmt.Sprintf("%v", data[key]), nil
	}

	return "", nil
}

func DeleteValue(key string) (string, error) {
	data, err := getContent()
	if err != nil {
		return "", err
	}

	if !existKey(data, key) {
		return "no_exist", nil
	}

	delete(data, key)

	if err := saveContent(data); err != nil {
		return "", err
	}

	return "success", nil
}

func GetKeys() ([]string, error) {
	var keys []string

	data, err := getContent()
	if err != nil {
		return keys, err
	}

	if len(data) > 0 {
		for key := range data {
			keys = append(keys, key)
		}

		return keys, nil
	}

	return keys, nil
}

func Clean() (bool, error) {
	data := map[string]interface{}{}

	if err := saveContent(data); err != nil {
		return false, err
	}

	return true, nil
}
