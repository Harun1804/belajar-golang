package main

import (
	"fmt"
	"example.com/note/helpers"
	"example.com/note/note"
)

func main() {
	title, content := getNoteData()
	_, err := note.New(title, content)
	helpers.DisplayError(err)
}

func getNoteData() (string, string) {
	title := getUserInput("Enter note title: ")
	content := getUserInput("Enter note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)

	return input
}