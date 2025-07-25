package services

import (
	"example.com/rest-api/models"
	"example.com/rest-api/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) CreateUser(u *models.User) error {
	return s.Repo.CreateUser(u)
}

func (s *UserService) Authenticate(u *models.User) error {
	return s.Repo.Authenticate(u)
}