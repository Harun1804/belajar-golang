package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/helpers"
	"example.com/note/note"
)

func main() {
	title, content := getNoteData()
	getNoteData, err := note.New(title, content)
	helpers.DisplayError(err)
	getNoteData.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Enter note title:")
	content := getUserInput("Enter note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%s ",prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	helpers.DisplayError(err)
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input
}