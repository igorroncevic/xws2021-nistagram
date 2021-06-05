package controllers

import (
	"context"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
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

func (s *PostGrpcController) CreatePost(ctx context.Context, in *protopb.Post) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var post *domain.Post
	post = post.ConvertFromGrpc(in)

	err := s.service.CreatePost(ctx, post)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create post")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (s *PostGrpcController) GetAllPosts(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts, err := s.service.GetAllPosts(ctx)

	if err != nil {
		return &protopb.ReducedPostArray{
			Posts: []*protopb.ReducedPost{},
		}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.ReducedPost{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.ReducedPostArray{
		Posts: responsePosts,
	}, nil
}

func (s *PostGrpcController) GetPostById(ctx context.Context, id string) (*protopb.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	post, err := s.service.GetPostById(ctx, id)

	if err != nil {
		return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
	}

	grpcPost := post.ConvertToGrpc()

	return grpcPost, nil
}

func (s *PostGrpcController) RemovePost(ctx context.Context, id string) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemovePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := s.service.RemovePost(ctx, id)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (s *PostGrpcController) SearchContentByLocation(ctx context.Context, in *protopb.SearchLocationRequest) (*protopb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var location string = in.Location

	posts, err := s.service.SearchContentByLocation(ctx, location)
	if err != nil {
		return &protopb.ReducedPostArray{
			Posts: []*protopb.ReducedPost{},
		}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.ReducedPost{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.ReducedPostArray{
		Posts: responsePosts,
	}, nil
}

func (s *PostGrpcController) GetPostsByHashtag(ctx context.Context, in *protopb.Hashtag) (*protopb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsByHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts, err := s.service.GetPostsByHashtag(ctx, in.Text)
	if err != nil {
		return &protopb.ReducedPostArray{
			Posts: []*protopb.ReducedPost{},
		}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.ReducedPost{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.ReducedPostArray{
		Posts: responsePosts,
	}, nil
}
