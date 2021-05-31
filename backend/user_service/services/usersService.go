package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/encryption"
	"gorm.io/gorm"
)

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(db *gorm.DB) (*UserService, error){
	repository, err := repositories.NewUserRepo(db)

	return &UserService{
		repository: repository,
	}, err
}

func (service *UserService) GetAllUsers() ([]persistence.User, error) {
	return service.repository.GetAllUsers()
}

func (service *UserService) CreateUser(user *persistence.User) error {
	user.Password = encryption.HashAndSalt([]byte(user.Password))
	return service.repository.CreateUser(user)
}

func (service *UserService) LoginUser(ctx context.Context, data common.Credentials) error {
	return nil //service.repository.CheckPassword(ctx, data)
}

func (service *UserService) UpdateUserProfile(userDTO domain.User) (bool, error) {
	if userDTO.Username == "" || userDTO.Email == "" {
		return false, errors.New("username or email can not be empty string")
	}

	return service.repository.UpdateUserProfile(userDTO)
}
