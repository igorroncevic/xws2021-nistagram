package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/proto"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	otgo "github.com/opentracing/opentracing-go"
	"io"
)

type Server struct {
	closer io.Closer
	tracer otgo.Tracer
	proto.UnimplementedFollowersServer
	followerController *FollowersGrpcController
}

func NewServer(driver neo4j.Driver) (*Server, error) {
	followerController, _ := NewFollowersController(driver)
	tracer, closer := tracer.Init("recommendationService")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		tracer: tracer,
		closer: closer,
		followerController: followerController,
	}, nil
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.closer
}

func (s *Server) CreateUserConnection(ctx context.Context, in *proto.CreateFollowerRequest) (*proto.EmptyResponse, error) {
	return s.followerController.CreateUserConnection(ctx, in)
}

func (s *Server) GetAllFollowers(ctx context.Context, in *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	return s.followerController.GetAllFollowers(ctx, in)
}

func (s *Server) GetAllFollowing(ctx context.Context,  in *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	return s.followerController.GetAllFollowing(ctx, in)
}

func (s *Server) DeleteBiDirectedConnection(ctx context.Context, in *proto.CreateFollowerRequest) (*proto.EmptyResponse, error) {
	return s.followerController.DeleteBiDirectedConnection(ctx, in)
}

func (s *Server) DeleteDirectedConnection(ctx context.Context, in *proto.CreateFollowerRequest) (*proto.EmptyResponse, error) {
	return s.followerController.DeleteDirectedConnection(ctx, in)
}

func (s *Server) CreateUser(ctx context.Context, in *proto.CreateUserRequest ) (*proto.EmptyResponse, error) {
	return s.followerController.CreateUser(ctx, in)
}

