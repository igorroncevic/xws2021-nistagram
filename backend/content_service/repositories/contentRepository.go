package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/content_service/util/images"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type ContentRepository interface {
	GetAllPosts(context.Context) ([]persistence.Post, error)
	CreatePost(context.Context, *domain.Post) error
	GetPostById(context.Context, string) (*persistence.Post, error)
}

type contentRepository struct {
	DB *gorm.DB
}

func NewContentRepo(db *gorm.DB) (*contentRepository, error) {
	if db == nil {
		panic("ContentRepository not created, gorm.DB is nil")
	}

	return &contentRepository{ DB: db }, nil
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

	postToSave := &persistence.Post{
		Id:          uuid.NewV4().String(),
		UserId:      post.UserId,
		IsAd:        post.IsAd,
		Type:        post.Type,
		Description: post.Description,
		Location:    post.Location,
		CreatedAt:   time.Now(),
	}

	result := repository.DB.Create(&postToSave)
	if result.Error != nil || result.RowsAffected != 1 {
		return errors.New("cannot save post")
	}

	savedMedia := []string{}
	for _, media := range post.Media{
		mimeType, err := images.GetImageType(media.Content)
		if err != nil{
			images.RemoveImages(savedMedia)
			return err
		}

		t := time.Now()
		formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
		name := post.UserId + "_" + formatted + "." + mimeType
		savedMedia = append(savedMedia, name)

		err = images.SaveImage(name, media.Content)
		if err != nil{
			images.RemoveImages(savedMedia)
			repository.DB.Delete(&postToSave)
			return errors.New("cannot save post")
		}

		media.PostId = postToSave.Id
		// TODO Save media to database
	}

	return nil
}