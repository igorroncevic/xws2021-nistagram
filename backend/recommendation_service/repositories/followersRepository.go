package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowersRepository interface {
	CreateUserConnection(context.Context, model.Follower) (bool, error)
	GetAllFollowers(context.Context, string) ([]model.User, error)
	GetAllFollowing(context.Context, string) ([]model.User, error)
	CreateUser(context.Context, model.User) (bool, error)
	DeleteDirectedConnection(context.Context, model.Follower) (bool, error)
	DeleteBiDirectedConnection (context.Context, model.Follower) (bool, error)
}

type followersRepository struct {
	driver neo4j.Driver
}

func NewFollowersRepository(driver neo4j.Driver) (*followersRepository, error) {
	if driver == nil {
		panic("FollowersRepository not created, driver is nil")
	}

	return &followersRepository{driver: driver}, nil
}

func (repository *followersRepository) GetUsers(ctx context.Context, userId string, query string) ([]model.User, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	var users []model.User
	_ , err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			query,
			map[string]interface{}{
				"UserId" : userId,
			})

		if err != nil {
			return nil, err
		}
		for result.Next() {
			user := result.Record().Values[0].(string)
			users = append(users, model.User{
				UserId: user,
			})

		}
		return users, nil
	})

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *followersRepository) GetAllFollowing(ctx context.Context, userId string) ([]model.User, error){
	query := "MATCH (a:User {id : $UserId})-[r:Follows]->(b:User) RETURN b.id"
	return repository.GetUsers(ctx, userId, query)
}

func (repository *followersRepository) GetAllFollowers(ctx context.Context, userId string) ([]model.User, error){
	query := "MATCH (b:User)-[r:Follows]->(a:User {id : $UserId}) RETURN b.id"
	return repository.GetUsers(ctx, userId, query)
}

func (repository *followersRepository) CreateUserConnection(ctx context.Context, f model.Follower) (bool, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result , err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (a:User {id : $UserId}), (b:User {id : $FollowerId}) CREATE" +
				"(a)-[:Follows {UserId : $UserId, FollowerId : $FollowerId ,IsMuted : $IsMuted," +
				" IsCloseFriend : $IsCloseFriend, IsApprovedRequest : $IsApprovedRequest, " +
				"IsNotificationEnabled : $IsNotificationEnabled}]->(b)" +
				"RETURN a",
				map[string]interface{}{
				"UserId" : f.UserId,
				"FollowerId" : f.FollowerId,
				"IsMuted" : f.IsMuted,
				"IsCloseFriend" : f.IsCloseFriends,
				"IsApprovedRequest" : f.IsApprovedRequest,
				"IsNotificationEnabled" : f.IsNotificationEnabled,
			})

		if err != nil {
			return nil, err
		}
		if result.Next()  {
			return result.Record().Values[0], nil
		}
		return false, errors.New("error: can not create follower connection")
	})
	if err != nil || result == nil {
		return false, err
	}
	return true, nil
}

func (repository *followersRepository) CreateUser(ctx context.Context, u model.User) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result , err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MERGE (n:User {id : $UserId}) RETURN n",//MERGE -> create if not exists, else match
			map[string]interface{}{
				"UserId" : u.UserId,
			})

		if err != nil {
			return nil, err
		}
		if result.Next()  {
			return true, nil
		}
		return false, errors.New("error: can not create user ")
	})
	if err != nil || result == nil {
		return false, err
	}
	return true, nil
}

func (repository *followersRepository) DeleteDirectedConnection(ctx context.Context, f model.Follower) (bool, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_ , err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run(
			"MATCH (a:User {id : $UserId})-[r:Follows]->(b:User {id : $FollowerId}) DELETE r",
			map[string]interface{}{
				"UserId" : f.UserId,
				"FollowerId" : f.FollowerId,
			})

		if err != nil {
			return false, err
		}
		return true, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil


}

//Kada se useri blokiraju - pozivao sam dva puta
func (repository *followersRepository) DeleteBiDirectedConnection(ctx context.Context, f model.Follower) (bool, error) {
	return false, nil
}