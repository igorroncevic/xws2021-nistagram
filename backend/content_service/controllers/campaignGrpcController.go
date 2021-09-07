package controllers

import (
	"context"
	"errors"
	"github.com/igorroncevic/xws2021-nistagram/common"
	"github.com/igorroncevic/xws2021-nistagram/common/kafka_util"
	protopb "github.com/igorroncevic/xws2021-nistagram/common/proto"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/content_service/model/domain"
	"github.com/igorroncevic/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"net/http"
)

type CampaignGrpcController struct {
	service    			*services.CampaignService
	jwtManager 			*common.JWTManager
	userEventsProducer  *kafka_util.KafkaProducer
	performanceProducer *kafka_util.KafkaProducer
}

func NewCampaignController(db *gorm.DB, jwtManager *common.JWTManager, userEventsProducer *kafka_util.KafkaProducer, performanceProducer *kafka_util.KafkaProducer) (*CampaignGrpcController, error) {
	service, err := services.NewCampaignService(db)
	if err != nil {
		return nil, err
	}

	return &CampaignGrpcController{
		service,
		jwtManager,
		userEventsProducer,
		performanceProducer,
	}, nil
}

func (controller *CampaignGrpcController) GetCampaign(ctx context.Context, in *protopb.RequestId) (*protopb.Campaign, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaign")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	campaign, err := controller.service.GetCampaign(ctx, in.Id)
	if err != nil {
		return &protopb.Campaign{}, err
	}

	return campaign.ConvertToGrpc(), nil
}

func (controller *CampaignGrpcController) GetCampaigns(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.CampaignArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaigns")
	defer span.Finish()
	claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	campaigns, err := controller.service.GetCampaigns(ctx, claims.UserId)
	if err != nil {
		return &protopb.CampaignArray{}, err
	}

	responseCampaigns := []*protopb.Campaign{}
	for _, campaign := range campaigns {
		responseCampaigns = append(responseCampaigns, campaign.ConvertToGrpc())
	}

	return &protopb.CampaignArray{
		Campaigns: responseCampaigns,
	}, nil
}

func (controller *CampaignGrpcController) CreateCampaign(ctx context.Context, in *protopb.Campaign) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCampaign")
	defer span.Finish()
	claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	campaign := domain.Campaign{}
	campaign = campaign.ConvertFromGrpc(in)

	err := controller.service.CreateCampaign(ctx, campaign)
	if err != nil {
		controller.performanceProducer.WritePerformanceMessage(kafka_util.ContentService, kafka_util.CreateCampaignFunction, kafka_util.GetPerformanceMessage(kafka_util.CreateCampaignFunction, false) + ", user: " + claims.Email, http.StatusInternalServerError)
		return &protopb.EmptyResponseContent{}, err
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (controller *CampaignGrpcController) UpdateCampaign(ctx context.Context, in *protopb.Campaign) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateCampaign")
	defer span.Finish()
	claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var campaign domain.Campaign
	campaign = campaign.ConvertFromGrpc(in)

	if campaign.AgentId != claims.UserId {
		return &protopb.EmptyResponseContent{}, errors.New("cant update other agent's campaign")
	}

	err := controller.service.UpdateCampaign(ctx, campaign)
	if err != nil {
		controller.performanceProducer.WritePerformanceMessage(kafka_util.ContentService, kafka_util.UpdateCampaignFunction, kafka_util.GetPerformanceMessage(kafka_util.UpdateCampaignFunction, false) + ", user: " + claims.Email, http.StatusInternalServerError)
		controller.userEventsProducer.WriteUserEventMessage(kafka_util.CampaignUpdate, claims.UserId, kafka_util.GetUserEventMessage(kafka_util.CampaignUpdate, false) + ", campaign id - " + campaign.Id)
		return &protopb.EmptyResponseContent{}, err
	}

	controller.userEventsProducer.WriteUserEventMessage(kafka_util.CampaignUpdate, claims.UserId, kafka_util.GetUserEventMessage(kafka_util.CampaignUpdate, true) + ", campaign id - " + campaign.Id)
	return &protopb.EmptyResponseContent{}, nil
}

func (controller *CampaignGrpcController) DeleteCampaign(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteCampaign")
	defer span.Finish()
	claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := controller.service.DeleteCampaign(ctx, in.Id)
	if err != nil {
		controller.userEventsProducer.WriteUserEventMessage(kafka_util.DeleteCampaign, claims.UserId, kafka_util.GetUserEventMessage(kafka_util.DeleteCampaign, false) + ", campaign id - " + in.Id)
		return &protopb.EmptyResponseContent{}, err
	}

	controller.userEventsProducer.WriteUserEventMessage(kafka_util.DeleteCampaign, claims.UserId, kafka_util.GetUserEventMessage(kafka_util.DeleteCampaign, true) + ", campaign id - " + in.Id)
	return &protopb.EmptyResponseContent{}, nil
}

func (controller *CampaignGrpcController) GetCampaignStats(ctx context.Context, in *protopb.RequestId) (*protopb.CampaignStats, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaignStats")
	defer span.Finish()
	claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	stats, err := controller.service.GetCampaignStatistics(ctx, claims.UserId, in.Id)
	if err != nil {
		return &protopb.CampaignStats{}, err
	}

	return stats.ConvertToGrpc(), nil
}

func (controller *CampaignGrpcController) CreateCampaignRequest(ctx context.Context, in *protopb.CampaignInfluencerRequest) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCampaignRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var campaignRequest *domain.CampaignInfluencerRequest
	campaignRequest = campaignRequest.ConvertFromGrpc(in)

	err := controller.service.CreateCampaignRequest(ctx, campaignRequest)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "Bad request")
	}
	return &protopb.EmptyResponseContent{}, nil
}

func (controller *CampaignGrpcController) UpdateCampaignRequest(ctx context.Context, in *protopb.CampaignInfluencerRequest) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateCampaignRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var campaignRequest *domain.CampaignInfluencerRequest
	campaignRequest = campaignRequest.ConvertFromGrpc(in)

	err := controller.service.UpdateCampaignRequest(ctx, campaignRequest)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "Bad request")
	}
	return &protopb.EmptyResponseContent{}, nil
}

func (controller *CampaignGrpcController) GetCampaignRequestsByAgent(ctx context.Context, in *protopb.CampaignInfluencerRequest) (*protopb.CampaignRequestArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaignRequestsByAgent")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	requests, err := controller.service.GetCampaignRequestsByAgent(ctx, in.AgentId)
	if err != nil {
		return &protopb.CampaignRequestArray{}, err
	}

	var requestList []*protopb.CampaignInfluencerRequest
	for _, request := range requests {
		requestList = append(requestList, request.ConvertToGrpc())
	}

	finalResponse := protopb.CampaignRequestArray{CampaignRequests: requestList}
	return &finalResponse, nil
}
