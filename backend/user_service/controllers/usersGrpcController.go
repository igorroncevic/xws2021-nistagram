package controllers

import (
	"context"
	userspb "github.com/david-drvar/xws2021-nistagram/user_service/proto"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
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
		service:  service,
	}, nil
}

func (s *UserGrpcController) CreateUser(ctx context.Context, in *userspb.CreateUserRequest) (*userspb.User, error) {
	return s.service.CreateUser(ctx, in)
}

func (s *UserGrpcController) GetAllUsers(ctx context.Context, in *userspb.EmptyRequest) (*userspb.UsersResponse, error) {
	return s.service.GetAllUsers(ctx)
}