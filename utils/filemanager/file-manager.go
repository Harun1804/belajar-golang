package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		file.Close()
		return nil, errors.New("Error opening file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Error reading file line: " + err.Error())
	}
	file.Close()

	return lines, nil
}

func WriteJson(directory string, filePath string, data interface{}) error {
	err := CreateDirectoryIfNotExists(directory)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("Error creating JSON file: " + err.Error())
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Error writing JSON to file: " + err.Error())
	}

	file.Close()
	return nil
}

func CreateDirectoryIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return errors.New("Error creating directory: " + err.Error())
		}
	}
	return nil
}
