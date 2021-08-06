package services

import (
	"context"
	"errors"
	"github.com/igorroncevic/xws2021-nistagram/common/grpc_common"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/content_service/model/domain"
	"github.com/igorroncevic/xws2021-nistagram/content_service/repositories"

	"gorm.io/gorm"
)

type LikeService struct {
	likeRepository    repositories.LikeRepository
	contentRepository repositories.PostRepository
	postService       *PostService
}

func NewLikeService(db *gorm.DB) (*LikeService, error) {
	likeRepository, err := repositories.NewLikeRepo(db)
	if err != nil {
		return nil, err
	}

	contentRepository, err := repositories.NewPostRepo(db)
	postService, _ := NewPostService(db)
	if err != nil {
		return nil, err
	}

	return &LikeService{
		likeRepository,
		contentRepository,
		postService,
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
	} else if post == nil {
		return likes, errors.New("cannot retrieve likes for non-existing post")
	}

	dbLikes, err := service.likeRepository.GetLikesForPost(ctx, postId, isLike)
	if err != nil {
		return likes, errors.New("error retrieving likes for post")
	}

	for _, like := range dbLikes {
		username, err := grpc_common.GetUsernameById(ctx, like.UserId)
		if err == nil {
			likes = append(likes, like.ConvertToDomain(username))
		}
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
	} else if post == nil {
		return errors.New("cannot like non-existing post")
	}

	err = service.likeRepository.CreateLike(ctx, like)
	if err != nil {
		return err
	}

	if like.IsLike {
		return grpc_common.CreateNotification(ctx, post.UserId, like.UserId, "Like", post.Id)
	}
	return grpc_common.CreateNotification(ctx, post.UserId, like.UserId, "Dislike", post.Id)
}

func (service *LikeService) GetUserLikedOrDislikedPosts(ctx context.Context, id string, isLike bool) ([]domain.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserLikedOrDislikedPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []domain.Post{}

	postIds, err := service.likeRepository.GetUserLikedOrDislikedPostIds(ctx, id, isLike)
	if err != nil {
		return posts, err
	}

	for _, postId := range postIds {
		domainPost, err := service.postService.GetPostById(ctx, postId)
		if err != nil {
			return []domain.Post{}, err
		}
		posts = append(posts, domainPost)
	}

	return posts, nil
}
