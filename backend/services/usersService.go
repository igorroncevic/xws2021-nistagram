package services

import (
	"fmt"
	"xws2021-nistagram/backend/models"
	"xws2021-nistagram/backend/repositories"
	"xws2021-nistagram/backend/util/encryption"
)

type UserService struct {
	Repository repositories.UserRepository
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	users, error := service.Repository.GetAllUsers()
	return users, error
}

func (service *UserService) CreateUser(user *models.User) error {
	fmt.Println(user.Password)
	user.Password = encryption.HashAndSalt([]byte(user.Password))
	fmt.Println(user.Password)
	error := service.Repository.CreateUser(user)
	return error
}
