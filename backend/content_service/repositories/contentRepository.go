package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/content_service/util/images"
	"gorm.io/gorm"
)

type ContentRepository interface {
	GetAllPosts(context.Context) ([]persistence.Post, error)
	CreatePost(context.Context, *domain.Post) error
	GetPostById(context.Context, string) (*persistence.Post, error)
	RemovePost(context.Context, string) error

	GetCollectionsPosts(context.Context, string) ([]persistence.Post, error)
}

type contentRepository struct {
	DB *gorm.DB
	mediaRepository MediaRepository
	tagRepository   TagRepository
}

func NewContentRepo(db *gorm.DB) (*contentRepository, error) {
	if db == nil {
		panic("ContentRepository not created, gorm.DB is nil")
	}

	mediaRepository, _ := NewMediaRepo(db)
	tagRepository, _ := NewTagRepo(db)

	return &contentRepository{
		DB: db,
		mediaRepository: mediaRepository,
		tagRepository: tagRepository,
	}, nil
}

func (repository *contentRepository) GetAllPosts(ctx context.Context) ([]persistence.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []persistence.Post{}
	result := repository.DB.Order("created_at desc").Find(&posts)
	if result.Error != nil {
		return posts, result.Error
	}

	return posts, nil
}

func (repository *contentRepository) GetPostById(ctx context.Context, id string) (*persistence.Post, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	post := &persistence.Post{}
	result := repository.DB.Where("id = ?", id).First(&post)
	if result.Error != nil {
		return post, result.Error
	}

	return post, nil
}


func (repository *contentRepository) CreatePost(ctx context.Context, post *domain.Post) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func (tx *gorm.DB) error {
		var postToSave persistence.Post
		postToSave = postToSave.ConvertToPersistence(*post)

		result := repository.DB.Create(&postToSave)
		if result.Error != nil || result.RowsAffected != 1 {
			return errors.New("cannot save post")
		}

		for _, media := range post.Media{
			media.PostId = postToSave.Id
			dbMedia, err := repository.mediaRepository.CreateMedia(ctx, media)

			if err != nil{
				return errors.New("cannot save post")
			}

			for _, tag := range media.Tags{
				tag.MediaId = dbMedia.Id
				err := repository.tagRepository.CreateTag(ctx, tag)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil { return err }
	return nil
}

func (repository *contentRepository) RemovePost(ctx context.Context, postId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemovePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func (tx *gorm.DB) error{
		post := &persistence.Post{ Id: postId }
		result := repository.DB.First(&post)

		if result.Error != nil || result.RowsAffected != 1 {
			return errors.New("cannot remove non-existing post")
		}

		postMedia, err := repository.mediaRepository.GetMediaForPost(ctx, post.Id)
		if err != nil {
			return errors.New("cannot retrieve post's media")
		}

		for _, media := range postMedia {
			err := tx.Transaction(func (tx *gorm.DB) error {
				mediaTags, err := repository.tagRepository.GetTagsForMedia(ctx, media.Id)
				if err != nil {
					return err
				}

				for _, tag := range mediaTags{
					var tagPers persistence.Tag
					tagPers.ConvertToPersistence(tag)

					err := repository.tagRepository.RemoveTag(ctx, tagPers)
					if err != nil {
						return err
					}
				}

				err = images.RemoveImages([]string{media.Filename})
				if err != nil {
					return errors.New("cannot remove media's images")
				}

				err = repository.mediaRepository.RemoveMedia(ctx, media.Id)
				if err != nil{
					return errors.New("cannot remove post's media")
				}

				return nil
			})
			if err != nil { return err }
		}
		return nil
	})

	if err != nil { return err }
	return nil
}

func (repository *contentRepository) GetCollectionsPosts(ctx context.Context, id string) ([]persistence.Post, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "RemovePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []persistence.Post{}
	result := repository.DB.Model(&persistence.Post{}).
		Joins("left join collections ON posts.id = collections.id").
		Where("collections.id = id").
		Find(&posts)

	if result.Error != nil {
		return posts, result.Error
	}

	return posts, nil
}