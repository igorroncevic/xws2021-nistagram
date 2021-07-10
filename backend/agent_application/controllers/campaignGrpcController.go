package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/domain"
	"github.com/david-drvar/xws2021-nistagram/agent_application/services"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CampaignGrpcController struct {
	service    *services.CampaignService
	jwtManager *common.JWTManager
	logger     *logger.Logger
}

func NewCampaignController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger) (*CampaignGrpcController, error) {
	service, err := services.NewCampaignService(db)
	if err != nil {
		return nil, err
	}

	return &CampaignGrpcController{
		service,
		jwtManager,
		logger,
	}, nil
}

func (controller *CampaignGrpcController) CreateCampaignReport(ctx context.Context, in *protopb.RequestIdAgent) (*protopb.EmptyResponseAgent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCampaignReport")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	//if claims.Role != "Agent"{
	//	return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown,"only agents can generate reports")
	//}

	campaignStats, err := grpc_common.GetCampaignStats(ctx, in.Id)
	if err != nil { return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, err.Error()) }

	var stats domain.CampaignStats
	stats = stats.ConvertFromGrpc(campaignStats)

	err = controller.service.CreateCampaignReport(ctx, stats)
	if err != nil { return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, err.Error()) }

	return &protopb.EmptyResponseAgent{}, nil
}
