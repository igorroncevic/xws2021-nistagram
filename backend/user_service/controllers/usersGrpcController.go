package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
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
	/*users, err := s.service.GetAllUsers(ctx)

		if err != nil {
			return &userspb.UsersResponse{
				Users: []*userspb.UsersDTO{},
			}, status.Errorf(codes.Unknown, "Could retrieve users")
		}

	<<<<<<< HEAD
		responseUsers := []*userspb.UsersDTO{}
		for _, user := range users{
	=======
		responseUsers := []*userspb.User{}
		for _, user := range users {
	>>>>>>> master
			responseUsers = append(responseUsers, user.ConvertToGrpc())
		}

		return &userspb.UsersResponse{
			Users: responseUsers,
	<<<<<<< HEAD
		}, nil*/
	return &userspb.UsersResponse{}, nil
}

func (s *UserGrpcController) UpdateUserProfile(ctx context.Context, in *userspb.CreateUserDTORequest) (*userspb.EmptyResponse, error) {
	var user domain.User

	user = user.ConvertFromGrpc(in.User)
	_, err := s.service.UpdateUserProfile(ctx, user)
	if err != nil {
		return &userspb.EmptyResponse{}, status.Errorf(codes.Unknown, "Could not create user")
	}

	return &userspb.EmptyResponse{}, nil
}

func (s *UserGrpcController) UpdateUserPassword(ctx context.Context, in *userspb.CreatePasswordRequest) (*userspb.EmptyResponse, error) {
	var password domain.Password

	password = password.ConvertFromGrpc(in.Password)
	_, err := s.service.UpdateUserPassword(ctx, password)
	if err != nil {
		return &userspb.EmptyResponse{}, status.Errorf(codes.InvalidArgument, "Could not create user")
	}

	return &userspb.EmptyResponse{}, nil
}

func (s *UserGrpcController) SearchUser(ctx context.Context, in *userspb.SearchUserDtoRequest) (*userspb.UsersResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SearchUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user = domain.User{}

	user = user.ConvertFromGrpc(in.User)
	users, err := s.service.SearchUsersByUsernameAndName(ctx, &user)
	if err != nil {
		return nil, err
	}

	var usersList []*userspb.UsersDTO
	for _, user := range users {
		usersList = append(usersList, user.ConvertToGrpc())
	}

	finalResponse := userspb.UsersResponse{
		Users: usersList,
	}

	return &finalResponse, nil
}
