package controllers

import (
	"context"
	"github.com/igorroncevic/xws2021-nistagram/common"
	"github.com/igorroncevic/xws2021-nistagram/common/grpc_common"
	protopb "github.com/igorroncevic/xws2021-nistagram/common/proto"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/content_service/model/domain"
	"github.com/igorroncevic/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CommentGrpcController struct {
	service     *services.CommentService
	postService *services.PostService
	jwtManager  *common.JWTManager
}

func NewCommentController(db *gorm.DB, jwtManager *common.JWTManager) (*CommentGrpcController, error) {
	service, err := services.NewCommentService(db)
	if err != nil {
		return nil, err
	}

	postService, err := services.NewPostService(db)
	if err != nil {
		return nil, err
	}

	return &CommentGrpcController{
		service,
		postService,
		jwtManager,
	}, nil
}

func (s *CommentGrpcController) CreateComment(ctx context.Context, in *protopb.Comment) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateComment")
	defer span.Finish()
	claims, err := s.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var comment *domain.Comment
	comment = comment.ConvertFromGrpc(in)

	if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unauthenticated, "cannot create comment")
	} else if comment.UserId != claims.UserId {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "cannot create comment for someone else")
	}

	err = s.service.CreateComment(ctx, comment)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create comment")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (s *CommentGrpcController) GetCommentsForPost(ctx context.Context, id string) (*protopb.CommentsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateComment")
	defer span.Finish()
	claims, _ := s.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	post, err := s.postService.GetPostById(ctx, id)
	if err != nil {
		return &protopb.CommentsArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	if claims.UserId == "" {
		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, post.UserId)
		if err != nil {
			return &protopb.CommentsArray{}, status.Errorf(codes.Unknown, err.Error())
		}
		if !isPublic {
			return &protopb.CommentsArray{}, status.Errorf(codes.Unknown, "this post is not public")
		}
	} else if claims.UserId != post.UserId {
		following, err := grpc_common.CheckFollowInteraction(ctx, post.UserId, claims.UserId)
		if err != nil {
			return &protopb.CommentsArray{}, status.Errorf(codes.Unknown, err.Error())
		}

		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, post.UserId)
		if err != nil {
			return &protopb.CommentsArray{}, status.Errorf(codes.Unknown, err.Error())
		}

		isBlocked, err := grpc_common.CheckIfBlocked(ctx, post.UserId, claims.UserId)
		if err != nil {
			return &protopb.CommentsArray{}, status.Errorf(codes.Unknown, err.Error())
		}

		if (!following.IsApprovedRequest && !isPublic) || isBlocked {
			return &protopb.CommentsArray{}, status.Errorf(codes.PermissionDenied, "cannot access this post")
		}
	}

	comments, err := s.service.GetCommentsForPost(ctx, id)
	if err != nil {
		return &protopb.CommentsArray{}, status.Errorf(codes.Unknown, "Could not retrieve comments")
	}

	responseComments := []*protopb.Comment{}
	for _, comment := range comments {
		responseComments = append(responseComments, comment.ConvertToGrpc())
	}

	return &protopb.CommentsArray{
		Comments: responseComments,
	}, nil
}
