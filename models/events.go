package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Store() error{
	query := `
	INSERT INTO events (name, description, location, date_time, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	events = append(events, e)

	defer stmt.Close()
	return nil
}

func GetEvents() []Event {
	return events
}

