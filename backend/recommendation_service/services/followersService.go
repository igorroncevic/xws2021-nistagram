package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/model"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/repositories"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowersService struct {
	repository repositories.FollowersRepository
}

func NewFollowersService(driver neo4j.Driver) (*FollowersService, error) {
	repository, err := repositories.NewFollowersRepository(driver)

	return &FollowersService{
		repository: repository,
	}, err
}

func (service *FollowersService) CreateUserConnection(ctx context.Context, follower model.Follower)  error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	follower.IsNotificationEnabled = true
	follower.IsCloseFriends = false
	follower.IsMuted = false

	privacy, err := grpc_common.CheckUserProfilePublic(ctx, follower.FollowerId)
	if err != nil {
		return err
	}
	if !privacy {
		follower.IsApprovedRequest = false
		follower.RequestIsPending=true
		res, err := service.repository.CreateUserConnection(ctx, follower)
		if err != nil || res == false {
			return err
		}
		return grpc_common.CreateNotification(ctx, follower.FollowerId, follower.UserId,  "FollowPrivate", "")
	}else {
		follower.IsApprovedRequest = true
		follower.RequestIsPending=false
		res, err := service.repository.CreateUserConnection(ctx, follower)
		if err != nil || res == false {
			return err
		}
		return grpc_common.CreateNotification(ctx, follower.FollowerId ,follower.UserId, "FollowPublic", "")
	}

}

func (service *FollowersService) GetAllFollowers(ctx context.Context, userId string) ([]model.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFollowers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetAllFollowers(ctx, userId)
}

func (service *FollowersService) GetAllFollowing(ctx context.Context, userId string) ([]model.User, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFollowing")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetAllFollowing(ctx, userId)
}

func (service *FollowersService) GetAllFollowingsForHomepage(ctx context.Context, userId string) ([]model.User, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFollowingsForHomepagePosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetAllFollowingsForHomepage(ctx, userId)
}

func (service *FollowersService) GetCloseFriends(ctx context.Context, userId string) ([]model.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCloseFriends")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetCloseFriends(ctx, userId)
}

func (service *FollowersService) CreateUser(ctx context.Context, u model.User) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.CreateUser(ctx, u);
}

func (service *FollowersService) DeleteDirectedConnection(ctx context.Context, f model.Follower) (bool, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.DeleteDirectedConnection(ctx, f)
}

func (service *FollowersService) DeleteBiDirectedConnection(ctx context.Context, f model.Follower) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.DeleteBiDirectedConnection(ctx, f)
}

func (service *FollowersService)  UpdateUserConnection(ctx context.Context, f model.Follower) (*model.Follower,error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.UpdateUserConnection(ctx, f)
}

func (service *FollowersService) GetFollowersConnection(ctx context.Context, f model.Follower) (*model.Follower, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetFollowersConnection(ctx, f)
}


