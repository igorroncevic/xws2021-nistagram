package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
	"time"
)

type CampaignService struct {
	campaignRepository    repositories.CampaignRepository
	adService  			  *AdService
}

func NewCampaignService(db *gorm.DB) (*CampaignService, error){
	campaignRepository, err := repositories.NewCampaignRepo(db)
	if err != nil {
		return nil, err
	}

	adService, err := NewAdService(db)
	if err != nil { return nil, err }

	return &CampaignService{
		campaignRepository,
		adService,
	}, err
}

// Retrieve only a list, not all the posts from the campaign
func (service *CampaignService) GetCampaigns(ctx context.Context, userId string) ([]domain.Campaign, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaigns")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCampaigns, err := service.campaignRepository.GetCampaigns(ctx, userId)
	if err != nil { return []domain.Campaign{}, err }

	campaigns := []domain.Campaign{}
	for _, dbCampaign := range dbCampaigns{
		category, err := service.adService.GetAdCategory(ctx, dbCampaign.AdCategoryId)
		if err != nil { return []domain.Campaign{}, err }

		campaigns = append(campaigns, dbCampaign.ConvertToDomain([]domain.Ad{}, category)) //Ads will be retrieved upon click
	}

	return campaigns, nil
}

// Only accessible by Agent, who gets all ads from the campaign, including influencer's ads
func (service *CampaignService) GetCampaign(ctx context.Context, campaignId string) (domain.Campaign, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCampaign, err := service.campaignRepository.GetCampaign(ctx, campaignId)
	if err != nil { return domain.Campaign{}, err }

	ads, err := service.adService.GetAdsFromCampaign(ctx, campaignId)
	if err != nil { return domain.Campaign{}, err }

	category, err := service.adService.GetAdCategory(ctx, dbCampaign.AdCategoryId)
	if err != nil { return domain.Campaign{}, err }

	return dbCampaign.ConvertToDomain(ads, category), nil
}

func (service *CampaignService) CreateCampaign(ctx context.Context, campaign domain.Campaign) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if campaign.Name == "" { return errors.New("no name provided") }
	if campaign.StartDate.Equal(time.Time{}) { return errors.New("no start date provided") }
	if campaign.EndDate.Equal(time.Time{}) { return errors.New("no start date provided") }
	if campaign.Category.Id == "" { return errors.New("no ad category") }
	if campaign.StartDate.After(campaign.EndDate) { return errors.New("start date cannot be after end date") }
	if campaign.IsOneTime && !campaign.StartDate.Equal(campaign.EndDate){ campaign.EndDate = campaign.StartDate/*.Add(24 * time.Hour)*/ }
	if len(campaign.Ads) == 0 { return errors.New("no ads provided") }

	return service.campaignRepository.CreateCampaign(ctx, campaign)
}

// Updates on !isOneTime campaigns need to be taken in consideration after 24hrs
func (service *CampaignService) UpdateCampaign(ctx context.Context, campaign domain.Campaign) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if campaign.Name == "" { return errors.New("no name provided") }
	if campaign.StartDate.Equal(time.Time{}) { return errors.New("no start date provided") }
	if campaign.EndDate.Equal(time.Time{}) { return errors.New("no end date provided") }
	if campaign.Category.Id == "" { return errors.New("no ad category") }
	if campaign.StartDate.After(campaign.EndDate) { return errors.New("start date cannot be after end date") }

	return service.campaignRepository.UpdateCampaign(ctx, campaign)
}

func (service *CampaignService) DeleteCampaign(ctx context.Context, campaignId string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.campaignRepository.DeleteCampaign(ctx, campaignId)
}