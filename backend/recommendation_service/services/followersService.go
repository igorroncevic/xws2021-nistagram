package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/model"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/repositories"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowersService struct {
	repository repositories.FollowersRepository
}

func NewFollowersService(driver *neo4j.Driver) (*FollowersService, error) {
	repository, err := repositories.NewFollowersRepository(driver)

	return &FollowersService{
		repository: repository,
	}, err
}

func (service *FollowersService) CreateUserConnection(ctx context.Context, follower model.Follower) (bool, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.CreateUserConnection(ctx, follower)
}
