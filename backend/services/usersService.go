package services

import (
	models2 "xws2021-nistagram/models"
	repositories2 "xws2021-nistagram/repositories"
)

type UserService struct {
	Repository repositories2.UserRepository
}

func (service *UserService) GetAllUsers() ([]models2.User, error) {
	users, error := service.Repository.GetAllUsers()
	return users, error
}

func (service *UserService) CreateUser(user *models2.User) error {
	error := service.Repository.CreateUser(user)
	return error
}
