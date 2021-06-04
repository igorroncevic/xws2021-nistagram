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

type PostGrpcController struct {
	service *services.PostService
}

func NewPostController(db *gorm.DB) (*PostGrpcController, error) {
	service, err := services.NewPostService(db)
	if err != nil {
		return nil, err
	}

	return &PostGrpcController{
		service: service,
	}, nil
}

func (s *PostGrpcController) CreatePost(ctx context.Context, in *contentpb.Post) (*contentpb.EmptyResponse, error) {
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

func (s *PostGrpcController) GetAllPosts(ctx context.Context, in *contentpb.EmptyRequest) (*contentpb.ReducedPostArray, error) {
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

func (s *PostGrpcController) GetPostById(ctx context.Context, id string) (*contentpb.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	post, err := s.service.GetPostById(ctx, id)

	if err != nil {
		return &contentpb.Post{}, status.Errorf(codes.Unknown, err.Error())
	}

	grpcPost := post.ConvertToGrpc()

	return grpcPost, nil
}

func (s *PostGrpcController) RemovePost(ctx context.Context, id string) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemovePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := s.service.RemovePost(ctx, id)

	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &contentpb.EmptyResponse{}, nil
}

func (s *PostGrpcController) SearchContentByLocation(ctx context.Context, in *contentpb.SearchLocationRequest) (*contentpb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var location string = in.Location

	posts, err := s.service.SearchContentByLocation(ctx, location)
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

func (s *PostGrpcController) GetPostsByHashtag(ctx context.Context, in *contentpb.Hashtag) (*contentpb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsByHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts, err := s.service.GetPostsByHashtag(ctx, in.Text)
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
