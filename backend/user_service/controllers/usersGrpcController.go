package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	userspb "github.com/david-drvar/xws2021-nistagram/user_service/proto"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserGrpcController struct {
	service *services.UserService
}

func NewUserController(db *gorm.DB) (*UserGrpcController, error) {
	service, err := services.NewUserService(db)
	if err != nil {
		return nil, err
	}

	return &UserGrpcController{
		service: service,
	}, nil
}

func (s *UserGrpcController) CreateUser(ctx context.Context, in *userspb.CreateUserRequest) (*userspb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user persistence.User = persistence.User{}
	var userAdditionalInfo persistence.UserAdditionalInfo = persistence.UserAdditionalInfo{}
	user = *user.ConvertFromGrpc(in.User)
	userAdditionalInfo = *userAdditionalInfo.ConvertFromGrpc(in.User)

	err := s.service.CreateUserWithAdditionalInfo(ctx, &user, &userAdditionalInfo)
	if err != nil {
		return &userspb.EmptyResponse{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &userspb.EmptyResponse{}, nil
}

func (s *UserGrpcController) GetAllUsers(ctx context.Context, in *userspb.EmptyRequest) (*userspb.UsersResponse, error) {
	users, err := s.service.GetAllUsers(ctx)

	if err != nil {
		return &userspb.UsersResponse{
			Users: []*userspb.User{},
		}, status.Errorf(codes.Unknown, "Could retrieve users")
	}

	responseUsers := []*userspb.User{}
	for _, user := range users {
		responseUsers = append(responseUsers, user.ConvertToGrpc())
	}

	return &userspb.UsersResponse{
		Users: responseUsers,
	}, nil
}
