package services

import (
	"context"
	"github.com/igorroncevic/xws2021-nistagram/common/kafka_util"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/model"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/repositories"
	"gorm.io/gorm"
)

type PerformanceService struct {
	performanceRepository		repositories.PerformanceRepository
}

func NewPerformanceService(db *gorm.DB) (*PerformanceService, error) {
	performanceRepository, _ := repositories.NewPerformanceRepo(db)

	return &PerformanceService{
		performanceRepository,
	}, nil
}

func (service *PerformanceService) GetAllStats(ctx context.Context) ([]kafka_util.PerformanceMessage, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllStats")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return []kafka_util.PerformanceMessage{}, nil
}

func (service *PerformanceService) SaveEntry(ctx context.Context, message model.PerformanceMessage) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveEntry")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.performanceRepository.SaveEntry(ctx, message)
}
