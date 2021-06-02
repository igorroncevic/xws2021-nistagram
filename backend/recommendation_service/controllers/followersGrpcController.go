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

func NewFollowersController(driver *neo4j.Driver) (*FollowersGrpcController, error){
	service, err := services.NewFollowersService(driver)
	if err != nil {
		return nil, err
	}

	return &FollowersGrpcController{
		service: service,
	}, nil
}

func (controller *FollowersGrpcController) CreateUserConnection(ctx context.Context, in proto.CreateFollowerRequest) (*proto.EmptyResponse, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	var follower = model.Follower{}
	follower = *follower.ConvertFromGrpc(in.Follower)
	result, err := controller.service.CreateUserConnection(ctx, follower)
	if !result || err != nil {
		return &proto.EmptyResponse{}, errors.New("Could not make follow!")
	}
	return &proto.EmptyResponse{}, nil
}
