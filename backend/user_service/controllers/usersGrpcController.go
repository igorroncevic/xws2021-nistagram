package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	userspb "github.com/david-drvar/xws2021-nistagram/user_service/proto"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	otgo "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"io"
)

type UserGrpcController struct {
	userspb.UnimplementedUsersServer
	service *services.UserService
	tracer otgo.Tracer
	closer io.Closer
}

func NewUserController(db *gorm.DB) (*UserGrpcController, error) {
	service, err := services.NewUserService(db)
	if err != nil {
		return nil, err
	}

	tracer, closer := tracer.Init("userController")
	otgo.SetGlobalTracer(tracer)
	return &UserGrpcController{
		service:  service,
		tracer: tracer,
		closer: closer,
	}, nil
}

func (s *UserGrpcController) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *UserGrpcController) GetCloser() io.Closer {
	return s.closer
}

func (s *UserGrpcController) CreateUser(ctx context.Context, in *userspb.CreateUserRequest) (*userspb.User, error) {
	return s.service.CreateUser(ctx, in)
}

func (s *UserGrpcController) GetAllUsers(ctx context.Context, in *userspb.EmptyRequest) (*userspb.UsersResponse, error) {
	return s.service.GetAllUsers(ctx)
}