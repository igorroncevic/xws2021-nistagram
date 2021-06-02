package controllers

import (
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

func NewServer(driver *neo4j.Driver) (*Server, error) {
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