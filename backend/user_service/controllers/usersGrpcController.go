package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserGrpcController struct {
	service   *services.UserService
	jwtManager *common.JWTManager
}

func NewUserController(db *gorm.DB, jwtManager *common.JWTManager) (*UserGrpcController, error) {
	service, err := services.NewUserService(db)
	if err != nil {
		return nil, err
	}

	return &UserGrpcController{
		service,
		jwtManager,
	}, nil
}

func (s *UserGrpcController) CreateUser(ctx context.Context, in *protopb.CreateUserRequest) (*protopb.UsersDTO, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user persistence.User
	var userAdditionalInfo persistence.UserAdditionalInfo
	user = *user.ConvertFromGrpc(in.User)
	userAdditionalInfo = *userAdditionalInfo.ConvertFromGrpc(in.User)

	userDomain, err := s.service.CreateUserWithAdditionalInfo(ctx, &user, &userAdditionalInfo)
	if err != nil {
		return &protopb.UsersDTO{}, status.Errorf(codes.Unknown, err.Error())
	}

	userProto := userDomain.ConvertToGrpc()

	return userProto, nil
}

func (s *UserGrpcController) GetAllUsers(ctx context.Context, in *protopb.EmptyRequest) (*protopb.UsersResponse, error) {
	return &protopb.UsersResponse{}, nil
}

func (s *UserGrpcController) UpdateUserProfile(ctx context.Context, in *protopb.CreateUserDTORequest) (*protopb.EmptyResponse, error) {
	var user domain.User

	user = user.ConvertFromGrpc(in.User)
	_, err := s.service.UpdateUserProfile(ctx, user)
	if err != nil {
		return &protopb.EmptyResponse{}, status.Errorf(codes.Unknown, "Could not create user")
	}

	return &protopb.EmptyResponse{}, nil
}

func (s *UserGrpcController) UpdateUserPassword(ctx context.Context, in *protopb.CreatePasswordRequest) (*protopb.EmptyResponse, error) {
	var password domain.Password

	password = password.ConvertFromGrpc(in.Password)
	_, err := s.service.UpdateUserPassword(ctx, password)
	if err != nil {
		return &protopb.EmptyResponse{}, status.Errorf(codes.InvalidArgument, "Could not create user")
	}

	return &protopb.EmptyResponse{}, nil
}

func (s *UserGrpcController) SearchUser(ctx context.Context, in *protopb.SearchUserDtoRequest) (*protopb.UsersResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SearchUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user domain.User
	user = user.ConvertFromGrpc(in.User)

	users, err := s.service.SearchUsersByUsernameAndName(ctx, &user)
	if err != nil {
		return nil, err
	}

	var usersList []*protopb.UsersDTO
	for _, user := range users {
		usersList = append(usersList, user.ConvertToGrpc())
	}

	finalResponse := protopb.UsersResponse{
		Users: usersList,
	}

	return &finalResponse, nil
}

func (s *UserGrpcController) GetUserById(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.UsersDTO, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "LoginUser")
	defer span.Finish()
	claims, err := s.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if claims.UserId == "" || in.Id == "" {
		return &protopb.UsersDTO{}, status.Errorf(codes.Unauthenticated, "cannot retrieve this user")
	}

	if claims.UserId != in.Id{
		following, err := grpc_common.CheckFollowInteraction(ctx, in.Id, claims.UserId)
		if err != nil {
			return &protopb.UsersDTO{}, status.Errorf(codes.Unknown, "cannot retrieve this user")
		}

		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.Id)
		if err != nil {
			return &protopb.UsersDTO{}, status.Errorf(codes.Unknown, err.Error())
		}

		isBlocked, err := grpc_common.CheckIfBlocked(ctx, in.Id, claims.UserId)
		if err != nil {
			return &protopb.UsersDTO{}, status.Errorf(codes.Unknown, err.Error())
		}

		// If used is blocked or his profile is private and did not approve your request
		if isBlocked || (!isPublic && !following.IsApprovedRequest ) {
			return &protopb.UsersDTO{}, status.Errorf(codes.Unknown, "cannot retrieve this user, no connection available")
		}
	}

	user, err := s.service.GetUser(ctx, in.Id)
	if err != nil{
		return &protopb.UsersDTO{}, err
	}

	userResponse := user.ConvertToGrpc()

	return userResponse, nil
}

func (s *UserGrpcController) GetUsernameById(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.UsersDTO, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsernameById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	username, err := s.service.GetUsername(ctx, in.Id)
	if err != nil{
		return &protopb.UsersDTO{}, err
	}

	userResponse := &protopb.UsersDTO{
		Username:     username,
	}

	return userResponse, nil
}


func (s *UserGrpcController) LoginUser(ctx context.Context, in *protopb.LoginRequest) (*protopb.LoginResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "LoginUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var request domain.LoginRequest
	request = request.ConvertFromGrpc(in)

	user, err := s.service.LoginUser(ctx, request)
	if err != nil{
		return &protopb.LoginResponse{}, err
	}

	token, err := s.jwtManager.GenerateJwt(user.Id, user.Role.String())
	if err != nil {
		return &protopb.LoginResponse{}, err
	}

	return &protopb.LoginResponse{
		AccessToken: token,
		UserId:      user.Id,
		Username:    user.Username,
		Role:        user.Role.String(),
	}, nil
}
