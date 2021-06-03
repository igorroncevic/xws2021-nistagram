package repositories

import (
	"context"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/content_service/util/images"
	"gorm.io/gorm"
	"time"
)

type MediaRepository interface {
	GetMediaForPost(context.Context, string) ([]persistence.Media, error)
	CreateMedia(context.Context, domain.Media) (persistence.Media, error)
	RemoveMedia(context.Context, string) error
}

type mediaRepository struct {
	DB *gorm.DB
}

func NewMediaRepo(db *gorm.DB) (*mediaRepository, error) {
	if db == nil {
		panic("MediaRepository not created, gorm.DB is nil")
	}

	return &mediaRepository{ DB: db }, nil
}

func (repository *mediaRepository) GetMediaForPost(ctx context.Context, postId string) ([]persistence.Media, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetMediaForPost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	medias := []persistence.Media{}
	result := repository.DB.Where("post_id = ?", postId).Order("order_num asc").Find(&medias)

	if result.Error != nil {
		return []persistence.Media{}, result.Error
	}

	return medias, nil
}

func (repository *mediaRepository) CreateMedia(ctx context.Context, media domain.Media) (persistence.Media, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateMedia")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	mimeType, err := images.GetImageType(media.Content)
	if err != nil{
		return persistence.Media{}, err
	}

	t := time.Now()
	formatted := fmt.Sprintf("%s%d%02d%02d%02d%02d%02d%02d", media.PostId, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	name := formatted + "." + mimeType

	err = images.SaveImage(name, media.Content)
	if err != nil{
		return persistence.Media{}, err
	}

	var dbMedia *persistence.Media
	dbMedia = dbMedia.ConvertToPersistence(media, name)

	result := repository.DB.Create(dbMedia)
	if result.RowsAffected != 1 || result.Error != nil {
		_ = images.RemoveImages([]string{name})
		return persistence.Media{}, result.Error
	}

	return *dbMedia, nil
}

func (repository *mediaRepository) RemoveMedia(ctx context.Context, id string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveMedia")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	mediaPers := persistence.Media{Id: id}
	result := repository.DB.Delete(&mediaPers)

	if result.Error != nil || result.RowsAffected != 1 {
		return result.Error
	}

	return nil
}
