package repositories

import (
	"context"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/model"
	"gorm.io/gorm"
)

type UserEventRepository interface {
	SaveEntry(context.Context, model.UserEventMessage) error
	GetUsersActivity(context.Context, string) ([]model.UserEventMessage, error)
}

type userEventRepository struct {
	DB *gorm.DB
}

func NewUserEventRepo(db *gorm.DB) (*userEventRepository, error) {
	if db == nil {
		panic("UserEventRepository not created, gorm.DB is nil")
	}

	return &userEventRepository { DB: db }, nil
}

func (repository *userEventRepository) SaveEntry(ctx context.Context, message model.UserEventMessage) error{
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

func (repository *userEventRepository) GetUsersActivity(ctx context.Context, id string) ([]model.UserEventMessage, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveEntry")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	events := []model.UserEventMessage{}
	err := repository.DB.Model(&model.UserEventMessage{}).Where("user_id = ?", id).Find(&events)
	if err.Error != nil { return []model.UserEventMessage{}, err.Error }

	return events, nil
}