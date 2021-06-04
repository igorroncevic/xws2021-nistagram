package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	recommendationpb "github.com/david-drvar/xws2021-nistagram/user_service/proto_intercommunication"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/encryption"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
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

func (service *UserService) GetUser(id string) (domain.User, error) {

	return domain.User{}, nil
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
	user, err := service.repository.CreateUserWithAdditionalInfo(ctx, user, userAdditionalInfo)
	if err != nil {
		return errors.New("Cannot create user!")
	}

	//todo create user node in graph database
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(":8095", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := recommendationpb.NewFollowersClient(conn)

	createUserRequest := recommendationpb.CreateUserRequest{
		User: &recommendationpb.User{
			UserId: user.Id,
		},
	}

	_, err = c.CreateUser(context.Background(), &createUserRequest)
	if err != nil {
		log.Fatalf("could not create node user: %s", err)
	}

	return nil
}

func (service *UserService) LoginUser(ctx context.Context, data common.Credentials) error {
	return nil //service.repository.CheckPassword(ctx, data)
}

func (service *UserService) UpdateUserProfile(ctx context.Context, userDTO domain.User) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserProfile")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if userDTO.Username == "" || userDTO.Email == "" {
		return false, errors.New("username or email can not be empty string")
	}

	return service.repository.UpdateUserProfile(ctx, userDTO)
}

func (service *UserService) UpdateUserPassword(ctx context.Context, password domain.Password) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserPassword")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if password.NewPassword != password.RepeatedPassword {
		return false, errors.New("Passwords do not match!")
	}

	_, err := service.repository.UpdateUserPassword(ctx, password)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (service *UserService) SearchUsersByUsernameAndName(ctx context.Context, user *domain.User) ([]domain.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.SearchUsersByUsernameAndName(ctx, user)
}
