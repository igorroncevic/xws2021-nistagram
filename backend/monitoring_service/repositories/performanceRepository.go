package repositories

import (
	"context"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/model"
	"gorm.io/gorm"
)

type PerformanceRepository interface {
	GetAllEntries(context.Context) ([]model.PerformanceMessage, error)
	SaveEntry(context.Context, model.PerformanceMessage) error
}

type performanceRepository struct {
	DB *gorm.DB
}

func NewPerformanceRepo(db *gorm.DB) (*performanceRepository, error) {
	if db == nil {
		panic("PerformanceRepository not created, gorm.DB is nil")
	}

	return &performanceRepository { DB: db }, nil
}

func (repository *performanceRepository) GetAllEntries(ctx context.Context) ([]model.PerformanceMessage, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllEntries")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	messages := []model.PerformanceMessage{}
	result := repository.DB.Find(&messages)
	if result.Error != nil { return []model.PerformanceMessage{}, result.Error }

	return messages, nil
}

func (repository *performanceRepository) SaveEntry(ctx context.Context, message model.PerformanceMessage) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveEntry")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		result := repository.DB.Create(message)
		if result.Error != nil || result.RowsAffected != 1 {
			return result.Error
		}

		return nil
	})

	return err
}