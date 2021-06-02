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

type LikeGrpcController struct {
	service *services.LikeService
}

func NewLikeController(db *gorm.DB) (*LikeGrpcController, error) {
	service, err := services.NewLikeService(db)
	if err != nil {
		return nil, err
	}

	return &LikeGrpcController{
		service,
	}, nil
}

func (s *LikeGrpcController) CreateLike(ctx context.Context, in *contentpb.Like) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateLike")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var like *domain.Like
	like = like.ConvertFromGrpc(in)

	err := s.service.CreateLike(ctx, *like)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not create like")
	}

	return &contentpb.EmptyResponse{}, nil
}

func (s *LikeGrpcController) GetLikesForPost(ctx context.Context, id string, isLike bool) (*contentpb.LikesArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetLikesForPost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	likes, err := s.service.GetLikesForPost(ctx, id, isLike)

	if err != nil{
		return &contentpb.LikesArray{
			Likes: []*contentpb.Like{},
		}, status.Errorf(codes.Unknown, "Could not retrieve likes")
	}

	responseLikes := []*contentpb.Like{}
	for _, like := range likes{
		responseLikes = append(responseLikes, like.ConvertToGrpc())
	}

	return &contentpb.LikesArray{
		Likes: responseLikes,
	}, nil
}
