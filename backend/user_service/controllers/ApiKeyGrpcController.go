package controllers

import (
	"context"
	"github.com/igorroncevic/xws2021-nistagram/common"
	"github.com/igorroncevic/xws2021-nistagram/common/kafka_util"
	"github.com/igorroncevic/xws2021-nistagram/common/logger"
	protopb "github.com/igorroncevic/xws2021-nistagram/common/proto"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/user_service/services"
	"gorm.io/gorm"
	"net/http"
)

type ApiTokenGrpcController struct {
	service             *services.ApiKeyService
	performanceProducer *kafka_util.KafkaProducer
	jwtManager          *common.JWTManager
	logger              *logger.Logger
}

func NewApiTokenGrpcController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger, performanceProducer *kafka_util.KafkaProducer) (*ApiTokenGrpcController, error) {
	service, err := services.NewApiTokenService(db)
	if err != nil {
		return nil, err
	}
	return &ApiTokenGrpcController{
		service,
		performanceProducer,
		jwtManager,
		logger,
	}, nil
}

func (controller *ApiTokenGrpcController) GenerateApiToken(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.ApiTokenResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GenerateApiToken")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result, err := controller.service.GenerateApiToken(ctx, in.Id)
	if err != nil {
		controller.performanceProducer.WritePerformanceMessage(kafka_util.UserService, kafka_util.GenerateApiTokenFunction, "Generate API token failed", http.StatusInternalServerError)
		return nil, err
	}
	return &protopb.ApiTokenResponse{Token: result}, nil

}

func (controller *ApiTokenGrpcController) GetKeyByUserId(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.ApiTokenResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetKeyByUserId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	token, err := controller.service.GetKeyByUserId(ctx, in.Id)
	return &protopb.ApiTokenResponse{Token: token}, err
}

func (controller *ApiTokenGrpcController) ValidateKey(ctx context.Context, in *protopb.ApiTokenResponse) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "ValidateKey")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := controller.service.ValidateKey(ctx, in.Token)
	return &protopb.EmptyResponse{}, err

}
