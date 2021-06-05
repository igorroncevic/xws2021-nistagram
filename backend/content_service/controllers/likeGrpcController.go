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

func (c *LikeGrpcController) CreateLike(ctx context.Context, in *protopb.Like) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateLike")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var like *domain.Like
	like = like.ConvertFromGrpc(in)

	err := c.service.CreateLike(ctx, *like)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create like")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *LikeGrpcController) GetLikesForPost(ctx context.Context, id string, isLike bool) (*protopb.LikesArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetLikesForPost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	likes, err := c.service.GetLikesForPost(ctx, id, isLike)

	if err != nil {
		return &protopb.LikesArray{
			Likes: []*protopb.Like{},
		}, status.Errorf(codes.Unknown, "Could not retrieve likes")
	}

	responseLikes := []*protopb.Like{}
	for _, like := range likes {
		responseLikes = append(responseLikes, like.ConvertToGrpc())
	}

	return &protopb.LikesArray{
		Likes: responseLikes,
	}, nil
}
