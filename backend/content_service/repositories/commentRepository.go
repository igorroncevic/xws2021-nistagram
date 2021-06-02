package repositories

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"gorm.io/gorm"
)

type CommentRepository interface {
	GetCommentsForPost(context.Context, string) ([]persistence.Comment, error)
	GetCommentsNumForPost(context.Context, string) (int, error)
	CreateComment(context.Context, domain.Comment) error
}

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepo(db *gorm.DB) (*commentRepository, error) {
	if db == nil {
		panic("CommentRepository not created, gorm.DB is nil")
	}

	return &commentRepository{ DB: db }, nil
}

func (repository *commentRepository) GetCommentsForPost(ctx context.Context, postId string) ([]persistence.Comment, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCommentsForPost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	comments := []persistence.Comment{}
	result := repository.DB.Order("created_at asc").Where("post_id = ?", postId).Find(&comments)

	if result.Error != nil {
		return comments, result.Error
	}

	return comments, nil
}

func (repository *commentRepository) GetCommentsNumForPost(ctx context.Context, postId string) (int, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCommentsNumForPost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var comments int64
	result := repository.DB.Model(&persistence.Comment{}).Where("post_id = ?", postId).Count(&comments)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(comments), nil
}

func (repository *commentRepository) CreateComment(ctx context.Context, comment domain.Comment) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateComment")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var commentPers *persistence.Comment
	commentPers = commentPers.ConvertToPersistence(comment)

	result := repository.DB.Create(commentPers)

	if result.Error != nil || result.RowsAffected != 1 {
		return result.Error
	}

	return nil
}
