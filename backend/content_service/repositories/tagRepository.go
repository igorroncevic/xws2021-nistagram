package repositories

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"gorm.io/gorm"
)

type TagRepository interface {
	GetTagsForMedia(context.Context, string) ([]domain.Tag, error)
	CreateTag(context.Context, domain.Tag) error
    RemoveTag(context.Context, persistence.Tag) error
}

type tagRepository struct {
	DB *gorm.DB
}

func NewTagRepo(db *gorm.DB) (*tagRepository, error) {
	if db == nil {
		panic("TagRepository not created, gorm.DB is nil")
	}

	return &tagRepository{ DB: db }, nil
}

func (repository *tagRepository) GetTagsForMedia(ctx context.Context, mediaId string) ([]domain.Tag, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetTagsForMedia")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	// TODO Get Username from other microservice
	tags := []domain.Tag{}
	result := repository.DB.Model(&domain.Tag{}).
		Select("tags.user_id, tags.media_id").
		Joins("left join media on media.id = tags.media_id").
		Where("media.id = ?", mediaId).Find(&tags)

	if result.Error != nil {
		return tags, result.Error
	}

	return tags, nil
}

func (repository *tagRepository) CreateTag(ctx context.Context, tag domain.Tag) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateTag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var tagPers *persistence.Tag
	tagPers = tagPers.ConvertToPersistence(tag)

	result := repository.DB.Create(tagPers)

	if result.Error != nil || result.RowsAffected != 1 {
		return result.Error
	}

	return nil
}

func (repository *tagRepository) RemoveTag(ctx context.Context, tag persistence.Tag) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveTag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repository.DB.Delete(&tag)

	if result.Error != nil || result.RowsAffected != 1 {
		return result.Error
	}

	return nil
}
