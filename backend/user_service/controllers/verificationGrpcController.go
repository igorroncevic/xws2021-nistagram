package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type VerificationGrpcController struct {
	service    *services.VerificationService
	jwtManager *common.JWTManager
	logger     *logger.Logger
}

func NewVerificationController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger) (*VerificationGrpcController, error) {
	service, err := services.NewVerificationService(db)
	if err != nil {
		return nil, err
	}

	return &VerificationGrpcController{
		service,
		jwtManager,
		logger,
	}, nil
}

func (s *VerificationGrpcController) SubmitVerificationRequest(ctx context.Context, in *protopb.VerificationRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SubmitVerificationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("SubmitVerificationRequest", "User verification request submit attempt: "+in.UserId, logger.Info)

	//todo - napravi konvertor u domenski ili persistence, napisi u user additional info category, dekodiraj sliku, upisi verification request
	var verificationRequest domain.VerificationRequest
	verificationRequest = verificationRequest.ConvertFromGrpc(in)

	err := s.service.CreateVerificationRequest(ctx, verificationRequest)
	if err != nil {
		return &protopb.EmptyResponse{}, status.Errorf(codes.Unknown, "Could not create verification request")
	}

	return &protopb.EmptyResponse{}, nil
}

func (s *VerificationGrpcController) GetPendingVerificationRequests(ctx context.Context, in *protopb.EmptyRequest) (*protopb.VerificationRequestsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPendingVerificationRequests")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("GetPendingVerificationRequests", "Get pending requests attempt", logger.Info)

	verificationRequests, err := s.service.GetPendingVerificationRequests(ctx)
	if err != nil {
		return &protopb.VerificationRequestsArray{}, status.Errorf(codes.Unknown, "Could not get pending verification requests")
	}

	responseVerificationRequests := []*protopb.VerificationRequest{}
	for _, verificationRequest := range verificationRequests {
		responseVerificationRequests = append(responseVerificationRequests, verificationRequest.ConvertToGrpc())
	}

	return &protopb.VerificationRequestsArray{
		VerificationRequests: responseVerificationRequests,
	}, nil
}
