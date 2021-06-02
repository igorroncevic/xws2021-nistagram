package repositories

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowersRepository interface {
	CreateUserConnection(context.Context, model.Follower) (bool, error)
}

type followersRepository struct {
	driver *neo4j.Driver
}

func NewFollowersRepository(driver *neo4j.Driver) (*followersRepository, error) {
	if driver == nil {
		panic("FollowersRepository not created, driver is nil")
	}

	return &followersRepository{driver: driver}, nil
}

func (repository *followersRepository)	CreateUserConnection(ctx context.Context, follower model.Follower) (bool, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)


	return true, nil
}

