package services

import (
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/encryption"
)

type UserService struct {
	Repository repositories.UserRepository
}

func (service *UserService) GetAllUsers() ([]persistence.User, error) {
	return service.Repository.GetAllUsers()
}

func (service *UserService) CreateUser(user *persistence.User) error {
	user.Password = encryption.HashAndSalt([]byte(user.Password))
	return service.Repository.CreateUser(user)
}

func (service *UserService) LoginUser(data common.Credentials) error {
	return service.Repository.CheckPassword(data)
}

func (service *UserService) UpdateUserProfile(userDTO domain.User) (bool, error) {
	if userDTO.Username == "" || userDTO.Email == "" {
		return false, errors.New("username or email can not be empty string")
	}

	return service.Repository.UpdateUserProfile(userDTO)
}
