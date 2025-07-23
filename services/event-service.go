package services

import (
	"example.com/rest-api/models"
	"example.com/rest-api/repositories"
)

type EventService struct{
	Repo *repositories.EventRepository
}

func (s *EventService) GetEvents() ([]models.Event, error) {
	return s.Repo.GetEvents()
}

func (s *EventService) GetEvent(id int64) (*models.Event, error) {
	return s.Repo.GetEvent(id)
}

func (s *EventService) CreateEvent(event *models.Event) error {
	return s.Repo.Store(event)
}

func (s *EventService) UpdateEvent(id int64, event *models.Event) error {
	event.ID = id
	return s.Repo.UpdateEvent(event)
}

func (s *EventService) DeleteEvent(event *models.Event) error {
    return s.Repo.DeleteEvent(event)
}