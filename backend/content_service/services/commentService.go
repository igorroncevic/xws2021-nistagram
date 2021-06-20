package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type CommentService struct {
	commentRepository repositories.CommentRepository
	contentRepository repositories.PostRepository
}

func NewCommentService(db *gorm.DB) (*CommentService, error){
	commentRepository, err := repositories.NewCommentRepo(db)
	if err != nil {
		return nil, err
	}

	contentRepository, err := repositories.NewPostRepo(db)
	if err != nil {
		return nil, err
	}

	return &CommentService{
		commentRepository,
		contentRepository,
	}, err
}

func (service CommentService) CreateComment(ctx context.Context, comment *domain.Comment) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateComment")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if comment.Content == "" {
		return errors.New("cannot create empty comment")
	}

	post, err := service.contentRepository.GetPostById(ctx, comment.PostId)
	if post.Id == "" || err != nil {
		return errors.New("post does not exist")
	}

	err = service.commentRepository.CreateComment(ctx, *comment)
	if err != nil {
		return err
	}

	users, err := grpc_common.GetUsersForNotificationEnabled(ctx, comment.UserId, "IsCommentNotificationEnabled")
	if err != nil {
		return errors.New("Could not create notification")
	}
	for _, u := range users.Users {
		if u.UserId == post.UserId {
			grpc_common.CreateNotification(ctx, post.UserId, comment.UserId, "Comment", post.Id)
			break
		}
	}

	return nil
}

func (service CommentService) GetCommentsForPost(ctx context.Context, id string) ([]domain.Comment, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCommentsForPost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	comments := []domain.Comment{}

	dbComments, err := service.commentRepository.GetCommentsForPost(ctx, id)
	if err != nil{
		return comments, err
	}

	for _, comment := range dbComments{
		username, err := grpc_common.GetUsernameById(ctx, comment.UserId)
		if err == nil {
			comments = append(comments, comment.ConvertToDomain(username))
		}
	}

	return comments, nil
}