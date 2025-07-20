package utils

import (
	"bufio"
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
