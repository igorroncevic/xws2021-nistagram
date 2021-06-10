package controllers

import (
	"context"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmailGrpcController struct {
	service *services.EmailService
}

func NewEmailController() (*EmailGrpcController, error) {
	service, err := services.NewEmailService()
	if err != nil {
		return nil, err
	}

	return &EmailGrpcController{
		service: service,
	}, nil
}

func (s *EmailGrpcController) SendEmail(ctx context.Context, in *protopb.SendEmailDtoRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SendMail")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := s.service.SendEmail(ctx,in.To)

	if err != nil {
		return &protopb.EmptyResponse{}, status.Errorf(codes.Unknown, "Could not send mail")
	}

	return &protopb.EmptyResponse{}, nil
}