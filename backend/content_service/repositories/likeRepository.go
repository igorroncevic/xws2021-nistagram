package repositories

import (
	"context"
	"errors"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/content_service/model/domain"
	"github.com/igorroncevic/xws2021-nistagram/content_service/model/persistence"
	"gorm.io/gorm"
)

type LikeRepository interface {
	GetLikesForPost(context.Context, string, bool) ([]persistence.Like, error)
	GetLikesNumForPost(context.Context, string, bool) (int, error)
	CreateLike(context.Context, domain.Like) error
	GetUserLikedOrDislikedPostIds(ctx context.Context, id string, isLike bool) ([]string, error)
}

type likeRepository struct {
	postRepository PostRepository
	DB             *gorm.DB
}

func NewLikeRepo(db *gorm.DB) (*likeRepository, error) {
	if db == nil {
		panic("LikeRepository not created, gorm.DB is nil")
	}

	postRepository, _ := NewPostRepo(db)

	return &likeRepository{
		DB:             db,
		postRepository: postRepository,
	}, nil
}

func (repository *likeRepository) GetLikesForPost(ctx context.Context, postId string, isLike bool) ([]persistence.Like, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetLikesForPost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	likes := []persistence.Like{}
	result := repository.DB.Where("post_id = ? AND is_like = ?", postId, isLike).Find(&likes)

	if result.Error != nil {
		return []persistence.Like{}, result.Error
	}

	return likes, nil
}

func (repository *likeRepository) GetLikesNumForPost(ctx context.Context, postId string, isLike bool) (int, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetLikesNumForPost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var likes int64
	result := repository.DB.Model(&persistence.Like{}).Where("post_id = ? AND is_like = ?", postId, isLike).Count(&likes)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(likes), nil
}

func (repository *likeRepository) CreateLike(ctx context.Context, like domain.Like) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateLike")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	// Check if user has already liked a post
	var checkLike persistence.Like
	result := repository.DB.Where("post_id = ? AND user_id = ?", like.PostId, like.UserId).First(&checkLike)

	if result.Error != nil && result.Error.Error() != "record not found" {
		return errors.New("unable to create like")
	}

	// The post has been liked/disliked
	if result.RowsAffected != 0 {
		// Swap like for dislike or vice-versa
		if checkLike.IsLike != like.IsLike {
			result := repository.DB.Model(&persistence.Like{}).Where("post_id = ? AND user_id = ?", checkLike.PostId, checkLike.UserId).Update("is_like", like.IsLike)

			if result.Error != nil || result.RowsAffected != 1 {
				return errors.New("unable to update like")
			}
		} else {
			// Remove like
			result := repository.DB.Where("post_id = ? AND user_id = ?", checkLike.PostId, checkLike.UserId).Delete(checkLike)

			if result.Error != nil || result.RowsAffected != 1 {
				return errors.New("unable to remove like")
			}
		}

	} else {
		var newLike *persistence.Like
		newLike = newLike.ConvertToPersistence(like)
		// Like does not exist, create a new one
		result := repository.DB.Create(newLike)

		if result.Error != nil || result.RowsAffected != 1 {
			return errors.New("unable to create new like after deletion")
		}
	}

	return nil
}

func (repository *likeRepository) GetUserLikedOrDislikedPostIds(ctx context.Context, id string, isLike bool) ([]string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserLikedOrDislikedPostIds")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	likes := []persistence.Like{}
	if isLike == true {
		result := repository.DB.Where("user_id = ? AND is_like = 'true'", id).Find(&likes)
		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		result := repository.DB.Where("user_id = ? AND is_like = 'false'", id).Find(&likes)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	var postIds []string
	for _, like := range likes {
		postIds = append(postIds, like.PostId)
	}

	return postIds, nil
}
