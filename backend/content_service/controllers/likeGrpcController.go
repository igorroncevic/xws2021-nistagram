package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type LikeGrpcController struct {
	service 	*services.LikeService
	postService *services.PostService
	jwtManager  *common.JWTManager
}

func NewLikeController(db *gorm.DB, jwtManager *common.JWTManager) (*LikeGrpcController, error) {
	service, err := services.NewLikeService(db)
	if err != nil {
		return nil, err
	}

	postService, err := services.NewPostService(db)
	if err != nil {
		return nil, err
	}

	return &LikeGrpcController{
		service,
		postService,
		jwtManager,
	}, nil
}

func (c *LikeGrpcController) CreateLike(ctx context.Context, in *protopb.Like) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateLike")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == ""{
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId != in.UserId {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "cannot create likes for other people")
	}

	var like *domain.Like
	like = like.ConvertFromGrpc(in)

	err = c.service.CreateLike(ctx, *like)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create like")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *LikeGrpcController) GetLikesForPost(ctx context.Context, id string, isLike bool) (*protopb.LikesArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetLikesForPost")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	post, err := c.postService.GetPostById(ctx, id)
	if err != nil {
		return &protopb.LikesArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	following, err := grpc_common.CheckFollowInteraction(ctx, post.UserId, claims.UserId)
	if err != nil { return &protopb.LikesArray{}, status.Errorf(codes.Unknown, err.Error()) }

	isPublic, err := grpc_common.CheckIfPublicProfile(ctx, post.UserId)
	if err != nil { return &protopb.LikesArray{}, status.Errorf(codes.Unknown, err.Error()) }

	isBlocked, err := grpc_common.CheckIfBlocked(ctx, post.UserId, claims.UserId)
	if err != nil { return &protopb.LikesArray{}, status.Errorf(codes.Unknown, err.Error()) }

	if (!following.IsApprovedRequest && !isPublic) || isBlocked {
		return &protopb.LikesArray{}, status.Errorf(codes.PermissionDenied, "cannot access this post")
	}

	likes, err := c.service.GetLikesForPost(ctx, id, isLike)
	if err != nil {
		return &protopb.LikesArray{}, status.Errorf(codes.Unknown, "could not retrieve likes")
	}

	responseLikes := []*protopb.Like{}
	for _, like := range likes {
		responseLikes = append(responseLikes, like.ConvertToGrpc())
	}

	return &protopb.LikesArray{
		Likes: responseLikes,
	}, nil
}
