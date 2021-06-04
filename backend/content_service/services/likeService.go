package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type LikeService struct {
	likeRepository repositories.LikeRepository
	contentRepository repositories.ContentRepository
}

func NewLikeService(db *gorm.DB) (*LikeService, error){
	likeRepository, err := repositories.NewLikeRepo(db)
	if err != nil {
		return nil, err
	}

	contentRepository, err := repositories.NewContentRepo(db)
	if err != nil {
		return nil, err
	}

	return &LikeService{
		likeRepository,
		contentRepository,
	}, err
}

func (service *LikeService) GetLikesForPost(ctx context.Context, postId string, isLike bool) ([]domain.Like, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllLikes")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	likes := []domain.Like{}

	post, err := service.contentRepository.GetPostById(ctx, postId)
	if err != nil {
		return likes, errors.New("error retrieving likes for post")
	}else if post == nil {
		return likes, errors.New("cannot retrieve likes for non-existing post")
	}

	dbLikes, err := service.likeRepository.GetLikesForPost(ctx, postId, isLike)
	if err != nil {
		return likes, errors.New("error retrieving likes for post")
	}

	for _, like := range dbLikes{
		likes = append(likes, like.ConvertToDomain())
	}

	return likes, nil
}

func (service *LikeService) CreateLike(ctx context.Context, like domain.Like) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateLike")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	post, err := service.contentRepository.GetPostById(ctx, like.PostId)
	if err != nil {
		return errors.New("error retrieving post to like")
	}else if post == nil {
		return errors.New("cannot like non-existing post")
	}

	return service.likeRepository.CreateLike(ctx, like)
}
