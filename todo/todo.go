package todo

import (
	"errors"
	"fmt"
	"os"
	"encoding/json"
)

type Todo struct {
	Text   string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("text cannot be empty")
	}

	return Todo{
		Text: text,
	}, nil
}

func (n Todo) Display() {
	fmt.Printf("Title: %s\n", n.Text)
}

func (n Todo) Save() error{
	fileName := "todo.json"

	jsonData, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}