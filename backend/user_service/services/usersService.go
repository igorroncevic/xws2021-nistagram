package services

import (
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/encryption"
)

type UserService struct {
	Repository repositories.UserRepository
}

func (service *UserService) GetAllUsers() ([]model.User, error) {
	return service.Repository.GetAllUsers()
}

func (service *UserService) CreateUser(user *model.User) error {
	user.Password = encryption.HashAndSalt([]byte(user.Password))
	return service.Repository.CreateUser(user)
}

func (service *UserService) LoginUser(data common.Credentials) error {
	return service.Repository.CheckPassword(data)
}
