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

func (s *CommentGrpcController) CreateComment(ctx context.Context, in *protopb.Comment) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateComment")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var comment *domain.Comment
	comment = comment.ConvertFromGrpc(in)

	err := s.service.CreateComment(ctx, comment)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create comment")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (s *CommentGrpcController) GetCommentsForPost(ctx context.Context, id string) (*protopb.CommentsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateComment")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	comments, err := s.service.GetCommentsForPost(ctx, id)

	if err != nil {
		return &protopb.CommentsArray{
			Comments: []*protopb.Comment{},
		}, status.Errorf(codes.Unknown, "Could not retrieve comments")
	}

	responseComments := []*protopb.Comment{}
	for _, comment := range comments {
		responseComments = append(responseComments, comment.ConvertToGrpc())
	}

	return &protopb.CommentsArray{
		Comments: responseComments,
	}, nil
}
