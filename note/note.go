package note

import (
	"errors"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" {
		return Note{}, errors.New("title cannot be empty")
	}

	if content == "" {
		return Note{}, errors.New("content cannot be empty")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}