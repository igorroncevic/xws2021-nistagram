package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	userspb "github.com/david-drvar/xws2021-nistagram/user_service/proto"
	otgo "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"io"
)

type Server struct {
	userspb.UnimplementedUsersServer
	userController *UserGrpcController
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(db *gorm.DB) (*Server, error) {
	controller, _ := NewUserController(db)
	tracer, closer := tracer.Init("userService")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		userController: controller,
		tracer: tracer,
		closer: closer,
	}, nil
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.closer
}

func (s *Server) CreateUser(ctx context.Context, in *userspb.CreateUserRequest) (*userspb.User, error) {
	return s.userController.CreateUser(ctx, in)
}

func (s *Server) GetAllUsers(ctx context.Context, in *userspb.EmptyRequest) (*userspb.UsersResponse, error) {
	return s.userController.GetAllUsers(ctx, in)
}

