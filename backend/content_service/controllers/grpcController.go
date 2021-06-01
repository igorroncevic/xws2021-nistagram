package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	contentpb "github.com/david-drvar/xws2021-nistagram/content_service/proto"
	otgo "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"io"
)

type Server struct {
	contentpb.UnimplementedContentServer
	contentController *ContentGrpcController
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(db *gorm.DB) (*Server, error) {
	controller, _ := NewContentController(db)
	tracer, closer := tracer.Init("contentService")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		contentController: controller,
		tracer: tracer,
		closer: closer,
	}, nil
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.closer
}

func (s *Server) CreatePost(ctx context.Context, in *contentpb.Post) (*contentpb.EmptyResponse, error) {
	return s.contentController.CreatePost(ctx, in)
}

func (s *Server) GetAllPosts(ctx context.Context, in *contentpb.EmptyRequest) (*contentpb.PostArray, error) {
	return s.contentController.GetAllPosts(ctx, in)
}

