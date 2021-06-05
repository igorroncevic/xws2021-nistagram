package controllers

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/model"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/proto"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/services"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowersGrpcController struct {
	service *services.FollowersService
}

func NewFollowersController(driver neo4j.Driver) (*FollowersGrpcController, error) {
	service, err := services.NewFollowersService(driver)
	if err != nil {
		return nil, err
	}

	return &FollowersGrpcController{
		service: service,
	}, nil
}

func (controller *FollowersGrpcController) CreateUserConnection(ctx context.Context, in *proto.CreateFollowerRequest) (*proto.EmptyResponseFollowers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	var follower = model.Follower{}
	follower = *follower.ConvertFromGrpc(in.Follower)
	result, err := controller.service.CreateUserConnection(ctx, follower)
	if !result || err != nil {
		return &proto.EmptyResponseFollowers{}, errors.New("Could not make follow!")
	}
	return &proto.EmptyResponseFollowers{}, nil
}
func (controller *FollowersGrpcController) GetAllFollowing(ctx context.Context, in *proto.CreateUserRequestFollowers) (*proto.CreateUserResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFollowing")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user = model.User{}
	user = *user.ConvertFromGrpc(in.User)
	users, err := controller.service.GetAllFollowing(ctx, user.UserId)
	if err != nil {
		return nil, errors.New("Could not get all followings")
	}

	return &proto.CreateUserResponse{Users: user.ConvertAllToGrpc(users)}, err
}

func (controller *FollowersGrpcController) GetAllFollowers(ctx context.Context, in *proto.CreateUserRequestFollowers) (*proto.CreateUserResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFollowers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user = model.User{}
	user = *user.ConvertFromGrpc(in.User)
	users, err := controller.service.GetAllFollowers(ctx, user.UserId)
	if err != nil {
		return nil, errors.New("Could not get all followers!")
	}

	return &proto.CreateUserResponse{Users: user.ConvertAllToGrpc(users)}, err
}
func (controller *FollowersGrpcController) DeleteBiDirectedConnection(ctx context.Context, in *proto.CreateFollowerRequest) (*proto.EmptyResponseFollowers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteBiDirectedConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var follower = model.Follower{}
	follower = *follower.ConvertFromGrpc(in.Follower)
	_, err := controller.service.DeleteBiDirectedConnection(ctx, follower)

	if err != nil {
		return &proto.EmptyResponseFollowers{}, errors.New("Could not delete bidirected connection!")
	}
	return &proto.EmptyResponseFollowers{}, nil
}

func (controller *FollowersGrpcController) DeleteDirectedConnection(ctx context.Context, in *proto.CreateFollowerRequest) (*proto.EmptyResponseFollowers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteDirectedConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var follower = model.Follower{}
	follower = *follower.ConvertFromGrpc(in.Follower)
	_, err := controller.service.DeleteDirectedConnection(ctx, follower)

	if err != nil {
		return &proto.EmptyResponseFollowers{}, errors.New("Could not delete directed connection!")
	}
	return &proto.EmptyResponseFollowers{}, nil
}

func (controller *FollowersGrpcController) CreateUser(ctx context.Context, in *proto.CreateUserRequestFollowers) (*proto.EmptyResponseFollowers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var user = model.User{}
	user = *user.ConvertFromGrpc(in.User)
	_, err := controller.service.CreateUser(ctx, user)

	if err != nil {
		return &proto.EmptyResponseFollowers{}, errors.New("Could not create User!")
	}
	return &proto.EmptyResponseFollowers{}, nil
}

func (controller *FollowersGrpcController) UpdateUserConnection(ctx context.Context, in *proto.CreateFollowerRequest) (*proto.CreateFollowerResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var follower = model.Follower{}
	follower = *follower.ConvertFromGrpc(in.Follower)
	result, err := controller.service.UpdateUserConnection(ctx, follower)

	if err != nil {
		return nil, errors.New("Could not update follower info!")
	}

	return &proto.CreateFollowerResponse{Follower: result.ConvertToGrpc()}, nil
}

func (controller *FollowersGrpcController) GetFollowersConnection(ctx context.Context, in *proto.CreateFollowerRequest) (*proto.CreateFollowerResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var follower = model.Follower{}
	follower = *follower.ConvertFromGrpc(in.Follower)
	result, err := controller.service.GetFollowersConnection(ctx, follower)

	if err != nil {
		return nil, errors.New("Could not update follower info!")
	}

	return &proto.CreateFollowerResponse{Follower: result.ConvertToGrpc()}, nil
}
