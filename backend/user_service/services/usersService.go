package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	userspb "github.com/david-drvar/xws2021-nistagram/user_service/proto"
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

func (service *UserService) GetAllUsers(ctx context.Context) (*userspb.UsersResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	responseUsers := []*userspb.User{}

	users, err := service.repository.GetAllUsers()
	if err != nil {
		return &userspb.UsersResponse{
			Users: responseUsers,
		}, err
	}

	for _, user := range users{
		responseUsers = append(responseUsers, user.ConvertToGrpc())
	}

	response := &userspb.UsersResponse{
		Users: responseUsers,
	}

	return response, nil
}

func (service *UserService) CreateUser(ctx context.Context, request *userspb.CreateUserRequest) (*userspb.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	request.User.Password = encryption.HashAndSalt([]byte(request.User.Password))

	err := service.repository.CreateUser(ctx, request.User)

	return request.User, err
}

func (service *UserService) LoginUser(ctx context.Context, data common.Credentials) error {
	return nil //service.repository.CheckPassword(ctx, data)
}
