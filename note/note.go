package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"encoding/json"
)

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" {
		return Note{}, errors.New("title cannot be empty")
	}

	if content == "" {
		return Note{}, errors.New("content cannot be empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("Title: %s\n", n.Title)
	fmt.Printf("Content: %s\n", n.Content)
}

func (n Note) Save() error{
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonData, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}