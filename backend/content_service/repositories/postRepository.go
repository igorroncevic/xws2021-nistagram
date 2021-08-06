package repositories

import (
	"context"
	"errors"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/content_service/model/domain"
	"github.com/igorroncevic/xws2021-nistagram/content_service/model/persistence"
	"github.com/igorroncevic/xws2021-nistagram/content_service/util/images"
	"gorm.io/gorm"
	"time"
)

type PostRepository interface {
	GetAllPosts(context.Context, []string) ([]persistence.Post, error)
	CreatePost(context.Context, *domain.Post) (persistence.Post, error)
	GetPostById(context.Context, string) (*persistence.Post, error)
	RemovePost(context.Context, string, string) error
	GetPostsByLocation(ctx context.Context, location string) ([]persistence.Post, error)
	GetCollectionsPosts(context.Context, string) ([]persistence.Post, error)
	GetPostsForUser(context.Context, string) ([]persistence.Post, error)
	UpdateCreatedAt(context.Context, string, time.Time) error
}

type postRepository struct {
	DB                  *gorm.DB
	mediaRepository     MediaRepository
	tagRepository       TagRepository
	hashtagRepository   HashtagRepository
	complaintRepository ComplaintRepository
}

func NewPostRepo(db *gorm.DB) (*postRepository, error) {
	if db == nil {
		panic("PostRepository not created, gorm.DB is nil")
	}

	mediaRepository, _ := NewMediaRepo(db)
	tagRepository, _ := NewTagRepo(db)
	hashtagRepository, _ := NewHashtagRepo(db)
	complaintRepository, _ := NewComplaintRepo(db)

	return &postRepository{
		DB:                  db,
		mediaRepository:     mediaRepository,
		tagRepository:       tagRepository,
		hashtagRepository:   hashtagRepository,
		complaintRepository: complaintRepository,
	}, nil
}

func (repository *postRepository) GetAllPosts(ctx context.Context, followings []string) ([]persistence.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []persistence.Post{}
	result := repository.DB.Order("created_at desc").
		Where("is_ad = false AND created_at <= ? AND user_id IN (?)", time.Now(), followings).Find(&posts)
	if result.Error != nil {
		return posts, result.Error
	}

	return posts, nil
}

func (repository *postRepository) GetPostsForUser(ctx context.Context, id string) ([]persistence.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []persistence.Post{}
	result := repository.DB.Order("created_at desc").Where("user_id = ? AND created_at <= ? AND is_ad = false", id, time.Now()).Find(&posts)
	if result.Error != nil {
		return posts, result.Error
	}

	return posts, nil
}

func (repository *postRepository) GetPostById(ctx context.Context, id string) (*persistence.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	post := &persistence.Post{}
	result := repository.DB.Where("id = ?", id).First(&post)
	if result.Error != nil {
		return post, result.Error
	}

	return post, nil
}

func (repository *postRepository) CreatePost(ctx context.Context, post *domain.Post) (persistence.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var postToSave persistence.Post
	postToSave = postToSave.ConvertToPersistence(*post)

	if postToSave.CreatedAt.IsZero() || (postToSave.CreatedAt.Year() == 1970 && postToSave.CreatedAt.Month() == 1 && postToSave.CreatedAt.Day() == 1) {
		postToSave.CreatedAt = time.Now()
	}
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		result := repository.DB.Create(&postToSave)
		if result.Error != nil || result.RowsAffected != 1 {
			return errors.New("cannot save post")
		}

		var finalHashtags []persistence.Hashtag
		//create hashtags if not exist
		for _, hashtag := range post.Hashtags {
			var domainHashtag *domain.Hashtag
			domainHashtag, _ = repository.hashtagRepository.GetHashtagByText(ctx, hashtag.Text)
			if domainHashtag.Id == "" {
				domainHashtag, _ = repository.hashtagRepository.CreateHashtag(ctx, hashtag.Text)
			}
			finalHashtags = append(finalHashtags, persistence.Hashtag{Id: domainHashtag.Id, Text: domainHashtag.Text})
		}

		//bind post with hashtags
		if len(post.Hashtags) != 0 {
			err := repository.hashtagRepository.BindPostWithHashtags(ctx, postToSave.Id, finalHashtags)
			if err != nil {
				return errors.New("cannot bind post with hashtags")
			}
		}

		for _, media := range post.Media {
			media.PostId = postToSave.Id
			dbMedia, err := repository.mediaRepository.CreateMedia(ctx, media)

			if err != nil {
				return errors.New("cannot save post")
			}

			for _, tag := range media.Tags {
				tag.MediaId = dbMedia.Id
				err := repository.tagRepository.CreateTag(ctx, tag)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		return persistence.Post{}, err
	}
	return postToSave, nil
}

func (repository *postRepository) RemovePost(ctx context.Context, postId string, userId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemovePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		post := &persistence.Post{Id: postId}
		if userId != "" {
			post.UserId = userId
		} // Removing post from campaign/ad repo won't contain userId
		result := repository.DB.First(&post)

		if result.Error != nil || result.RowsAffected != 1 {
			return errors.New("cannot remove non-existing post")
		}

		result = repository.DB.Delete(&post)
		if result.Error != nil || result.RowsAffected != 1 {
			return errors.New("cannot remove post")
		}

		err := repository.complaintRepository.DeleteByPostId(ctx, postId)
		if err != nil {
			return err
		}
		postMedia, err := repository.mediaRepository.GetMediaForPost(ctx, post.Id)
		if err != nil {
			return errors.New("cannot retrieve post's media")
		}

		for _, media := range postMedia {
			err := tx.Transaction(func(tx *gorm.DB) error {
				mediaTags, err := repository.tagRepository.GetTagsForMedia(ctx, media.Id)
				if err != nil {
					return err
				}

				for _, tag := range mediaTags {
					var tagPers *persistence.Tag
					tagPers = tagPers.ConvertToPersistence(tag)

					err := repository.tagRepository.RemoveTag(ctx, *tagPers)
					if err != nil {
						return err
					}
				}

				err = images.RemoveImages([]string{media.Filename})
				if err != nil {
					return errors.New("cannot remove media's images")
				}

				err = repository.mediaRepository.RemoveMedia(ctx, media.Id)
				if err != nil {
					return errors.New("cannot remove post's media")
				}

				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (repository *postRepository) GetCollectionsPosts(ctx context.Context, id string) ([]persistence.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCollectionsPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []persistence.Post{}
	result := repository.DB.Model(&persistence.Post{}).
		Joins("left join favorites   ON posts.id = favorites.post_id").
		Joins("left join collections ON favorites.collection_id = collections.id").
		Where("collections.id = ? A", id).
		Find(&posts)

	if result.Error != nil {
		return posts, result.Error
	}

	return posts, nil
}

func (repository *postRepository) GetPostsByLocation(ctx context.Context, location string) ([]persistence.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsByLocation")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var posts []persistence.Post
	result := repository.DB.Where("location = ?", location).Find(&posts)
	if result.Error != nil {
		return posts, result.Error
	}

	return posts, nil
}

func (repository *postRepository) UpdateCreatedAt(ctx context.Context, id string, createdAt time.Time) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateCreatedAt")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repository.DB.Model(&persistence.Post{}).Where("id = ?", id).Update("created_at", createdAt)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
