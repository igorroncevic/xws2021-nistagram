package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type ContentService struct {
	repository repositories.ContentRepository
}

func NewContentService(db *gorm.DB) (*ContentService, error){
	repository, err := repositories.NewContentRepo(db)

	return &ContentService{
		repository: repository,
	}, err
}

func (service *ContentService) GetAllPosts(ctx context.Context) ([]domain.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	posts := []domain.Post{}

	dbPosts, err := service.repository.GetAllPosts(ctx)
	if err != nil{
		return posts, err
	}

	for _, post := range dbPosts{
		comments := []domain.Comment{}
		likes := []domain.Like{}
		dislikes := []domain.Like{}
		tags := []domain.Tag{}
		media := []domain.Media{}
		posts = append(posts, post.ConvertToDomain(comments, likes, dislikes, tags, media))
	}

	return posts, nil
}

func (service *ContentService) CreatePost(ctx context.Context, post *domain.Post) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if len(post.Media) == 0 {
		return errors.New("cannot create empty post")
	}

	return service.repository.CreatePost(ctx, post)
}
