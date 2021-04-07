package services

import (
	"xws2021-nistagram/models"
	"xws2021-nistagram/repositories"
)

type UserService struct {
	Repository repositories.UserRepository
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	users, error := service.Repository.GetAllUsers()
	return users, error
}

func (service *UserService) CreateUser(user *models.User) error {
	error := service.Repository.CreateUser(user)
	return error
}
