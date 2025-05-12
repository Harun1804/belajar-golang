package main

import (
	"errors"
	"fmt"
)

func main() {
	_, _, err := getNoteData()
	displayError(err)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Enter note title: ")
	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Enter note content: ")
	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	if input == "" { 
		return "", errors.New("input cannot be empty")
	}

	return input, nil
}

func displayError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}