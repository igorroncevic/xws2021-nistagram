package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	contentpb "github.com/david-drvar/xws2021-nistagram/content_service/proto"
	"github.com/david-drvar/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ContentGrpcController struct {
	service *services.ContentService
}

func NewContentController(db *gorm.DB) (*ContentGrpcController, error) {
	service, err := services.NewContentService(db)
	if err != nil {
		return nil, err
	}

	return &ContentGrpcController{
		service: service,
	}, nil
}

func (s *ContentGrpcController) CreatePost(ctx context.Context, in *contentpb.Post) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var post *domain.Post
	post = post.ConvertFromGrpc(in)

	err := s.service.CreatePost(ctx, post)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not create post")
	}

	return &contentpb.EmptyResponse{}, nil
}

func (s *ContentGrpcController) GetAllPosts(ctx context.Context, in *contentpb.EmptyRequest) (*contentpb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts, err := s.service.GetAllPosts(ctx)

	if err != nil {
		return &contentpb.ReducedPostArray{
			Posts: []*contentpb.ReducedPost{},
		}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*contentpb.ReducedPost{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &contentpb.ReducedPostArray{
		Posts: responsePosts,
	}, nil
}

func (s *ContentGrpcController) SearchContentByLocation(ctx context.Context, in *contentpb.SearchLocationRequest) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var location string = in.Location

	err, _ := s.service.SearchContentByLocation(ctx, location)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not create post")
	}

	return &contentpb.EmptyResponse{}, nil
}
