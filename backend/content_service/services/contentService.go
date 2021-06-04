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
	contentRepository repositories.ContentRepository
	commentRepository repositories.CommentRepository
	likeRepository 	  repositories.LikeRepository
	mediaRepository   repositories.MediaRepository
	tagRepository	  repositories.TagRepository
}

func NewContentService(db *gorm.DB) (*ContentService, error){
	contentRepository, err := repositories.NewContentRepo(db)
	if err != nil {
		return nil, err
	}

	commentRepository, err := repositories.NewCommentRepo(db)
	if err != nil {
		return nil, err
	}

	likeRepository, err := repositories.NewLikeRepo(db)
	if err != nil {
		return nil, err
	}

	mediaRepository, err := repositories.NewMediaRepo(db)
	if err != nil {
		return nil, err
	}

	tagRepository, err := repositories.NewTagRepo(db)
	if err != nil {
		return nil, err
	}

	return &ContentService{
		contentRepository,
		commentRepository,
		likeRepository,
		mediaRepository,
		tagRepository,
	}, err
}

func (service *ContentService) GetAllPosts(ctx context.Context) ([]domain.ReducedPost, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []domain.ReducedPost{}

	dbPosts, err := service.contentRepository.GetAllPosts(ctx)
	if err != nil{
		return posts, err
	}

	// TODO Retrieve all domain data
	for _, post := range dbPosts{
		commentsNum, err := service.commentRepository.GetCommentsNumForPost(ctx, post.Id)
		if err != nil{
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts comments")
		}

		likes, err := service.likeRepository.GetLikesNumForPost(ctx, post.Id, true)
		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts likes")
		}

		dislikes, err := service.likeRepository.GetLikesNumForPost(ctx, post.Id, false)
		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts dislikes")
		}

		media, err := service.mediaRepository.GetMediaForPost(ctx, post.Id)
		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts media")
		}

		convertedMedia := []domain.Media{}
		for _, single := range media{
			tags, err := service.tagRepository.GetTagsForMedia(ctx, single.Id)
			if err != nil {
				return []domain.ReducedPost{}, errors.New("unable to retrieve media tags")
			}

			converted, err := single.ConvertToDomain(tags)
			if err != nil {
				return []domain.ReducedPost{}, errors.New("unable to convert media")
			}

			convertedMedia = append(convertedMedia, converted)
		}

		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to convert posts media")
		}

		posts = append(posts, post.ConvertToDomainReduced(commentsNum, likes, dislikes, convertedMedia))
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

	return service.contentRepository.CreatePost(ctx, post)
}
