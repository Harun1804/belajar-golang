package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath string
	OutputDirectory string
	OutputFilePath string
}

func (fm FileManager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("Error opening file: " + err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.New("Error reading file line: " + err.Error())
	}

	return lines, nil
}

func (fm FileManager) WriteJson(data interface{}) error {
	err := createDirectoryIfNotExists(fm.OutputDirectory)
	if err != nil {
		return err
	}

	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Error creating JSON file: " + err.Error())
	}

	defer file.Close()
	time.Sleep(3 * time.Second) // Simulate a delay for writing

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Error writing JSON to file: " + err.Error())
	}

	return nil
}

func createDirectoryIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return errors.New("Error creating directory: " + err.Error())
		}
	}
	return nil
}

func New(inputFilePath, outputDirectory, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputDirectory: outputDirectory,
		OutputFilePath: outputFilePath,
	}
}
