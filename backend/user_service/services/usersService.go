package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/encryption"
	"gorm.io/gorm"
)

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(db *gorm.DB) (*UserService, error) {
	repository, err := repositories.NewUserRepo(db)

	return &UserService{
		repository: repository,
	}, err
}

func (service *UserService) GetAllUsers(ctx context.Context) ([]persistence.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetAllUsers(ctx)
}

func (service *UserService) CreateUser(ctx context.Context, user *persistence.User) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	user.Password = encryption.HashAndSalt([]byte(user.Password))
	return service.repository.CreateUser(ctx, user)
}

func (service *UserService) CreateUserWithAdditionalInfo(ctx context.Context, user *persistence.User, userAdditionalInfo *persistence.UserAdditionalInfo) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	user.Password = encryption.HashAndSalt([]byte(user.Password))
	return service.repository.CreateUserWithAdditionalInfo(ctx, user, userAdditionalInfo)
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
