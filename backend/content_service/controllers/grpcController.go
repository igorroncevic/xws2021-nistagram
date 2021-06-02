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
	commentController *CommentGrpcController
	likeController	  *LikeGrpcController
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(db *gorm.DB) (*Server, error) {
	contentController, _ := NewContentController(db)
	commentController, _ := NewCommentController(db)
	likeController, _ := NewLikeController(db)
	tracer, closer := tracer.Init("global_ContentGrpcController")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		contentController: contentController,
		commentController: commentController,
		likeController: likeController,
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

func (s *Server) GetAllPosts(ctx context.Context, in *contentpb.EmptyRequest) (*contentpb.ReducedPostArray, error) {
	return s.contentController.GetAllPosts(ctx, in)
}

func (s *Server) CreateComment(ctx context.Context, in *contentpb.Comment) (*contentpb.EmptyResponse, error) {
	return s.commentController.CreateComment(ctx, in)
}

func (s *Server) GetCommentsForPost(ctx context.Context, in *contentpb.RequestId) (*contentpb.CommentsArray, error) {
	return s.commentController.GetCommentsForPost(ctx, in.Id)
}

func (s *Server) CreateLike(ctx context.Context, in *contentpb.Like) (*contentpb.EmptyResponse, error) {
	return s.likeController.CreateLike(ctx, in)
}

func (s *Server) GetLikesForPost(ctx context.Context, in *contentpb.LikeRequest) (*contentpb.LikesArray, error) {
	return s.likeController.GetLikesForPost(ctx, in)
}


