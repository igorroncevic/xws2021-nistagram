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

type CommentGrpcController struct {
	service *services.CommentService
}

func NewCommentController(db *gorm.DB) (*CommentGrpcController, error) {
	service, err := services.NewCommentService(db)
	if err != nil {
		return nil, err
	}

	return &CommentGrpcController{
		service,
	}, nil
}

func (s *CommentGrpcController) CreateComment(ctx context.Context, in *contentpb.Comment) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateComment")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var comment *domain.Comment
	comment = comment.ConvertFromGrpc(in)

	err := s.service.CreateComment(ctx, comment)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not create comment")
	}

	return &contentpb.EmptyResponse{}, nil
}

func (s *CommentGrpcController) GetCommentsForPost(ctx context.Context, id string) (*contentpb.CommentsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateComment")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	comments, err := s.service.GetCommentsForPost(ctx, id)

	if err != nil{
		return &contentpb.CommentsArray{
			Comments: []*contentpb.Comment{},
		}, status.Errorf(codes.Unknown, "Could not retrieve comments")
	}

	responseComments := []*contentpb.Comment{}
	for _, comment := range comments{
		responseComments = append(responseComments, comment.ConvertToGrpc())
	}

	return &contentpb.CommentsArray{
		Comments: responseComments,
	}, nil
}