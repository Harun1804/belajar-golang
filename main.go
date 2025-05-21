package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/helpers"
	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputer interface {
	saver
	displayer
	// DoSomething()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("hello")
	title, content := getNoteData()
	getNoteData, err := note.New(title, content)

	helpers.DisplayError(err)
	err = outputDate(getNoteData)
	helpers.DisplayError(err)
	fmt.Println("Note saved successfully!")

	text := getTodoData()
	getTodoData, err := todo.New(text)
	helpers.DisplayError(err)
	outputDate(getTodoData)
	// helpers.DisplayError(err)
	fmt.Println("Todo saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter note title:")
	content := getUserInput("Enter note content:")

	return title, content
}

func getTodoData() (string) {
	text := getUserInput("Enter todo text:")

	return text
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

func saveData(s saver) error {
	err := s.Save()
	if err != nil {
		return err
	}

	return nil
}

func outputDate(output outputer) error {
	output.Display()
	return saveData(output)
}

func printSomething(value interface{}) {
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("string: ", stringVal)
		return
	}

	intVal, ok := value.(int)
	if ok {
		fmt.Println("int: ", intVal)
		return
	}

	// switch value.(type) {
	// 	case int:
	// 		fmt.Println("integer: ",value)
	// 	case string:
	// 		fmt.Println("string: ",value)
	// 	case float64:
	// 		fmt.Println("float: ",value)
	// 	default: 
	// 		fmt.Println("unknown type: ",value)
	// }
}