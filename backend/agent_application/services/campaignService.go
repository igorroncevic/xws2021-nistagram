package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/domain"
	"github.com/david-drvar/xws2021-nistagram/agent_application/repositories"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"gorm.io/gorm"
)

type CampaignService struct {
	campaignRepository repositories.CampaignRepository
}

func NewCampaignService(db *gorm.DB) (*CampaignService, error) {
	campaignRepository, err := repositories.NewCampaignRepo(db)

	return &CampaignService{
		campaignRepository,
	}, err
}

func (service *CampaignService) CreateCampaignReport(ctx context.Context, stats domain.CampaignStats) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCampaignReport")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.campaignRepository.CreateCampaignReport(ctx, stats)
}