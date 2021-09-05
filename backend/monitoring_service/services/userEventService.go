package services

import (
	"context"
	"github.com/igorroncevic/xws2021-nistagram/common/kafka_util"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/model"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/repositories"
	"gorm.io/gorm"
)

type UserEventService struct {
	userEventRepository		repositories.UserEventRepository
}

func NewUserEventService(db *gorm.DB) (*UserEventService, error) {
	userEventRepository, _ := repositories.NewUserEventRepo(db)

	return &UserEventService{
		userEventRepository,
	}, nil
}

func (service *UserEventService) GetAllStats(ctx context.Context) ([]kafka_util.UserEventMessage, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllStats")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return []kafka_util.UserEventMessage{}, nil
}

func (service *UserEventService) SaveEntry(ctx context.Context, message model.UserEventMessage) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveEntry")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.userEventRepository.SaveEntry(ctx, message)
}

func (service *UserEventService) GetUsersActivity(ctx context.Context, id string) ([]model.UserEventMessage, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveEntry")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.userEventRepository.GetUsersActivity(ctx, id)
}