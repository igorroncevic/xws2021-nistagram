package services

import (
	"xws2021-nistagram/backend/models"
	"xws2021-nistagram/backend/models/dtos"
	"xws2021-nistagram/backend/repositories"
	"xws2021-nistagram/backend/util/encryption"
)

type UserService struct {
	Repository repositories.UserRepository
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	return service.Repository.GetAllUsers()
}

func (service *UserService) CreateUser(user *models.User) error {
	user.Password = encryption.HashAndSalt([]byte(user.Password))
	return service.Repository.CreateUser(user)
}

func (service *UserService) LoginUser(data dtos.LoginDTO) error {
	return service.Repository.CheckPassword(data)
}
