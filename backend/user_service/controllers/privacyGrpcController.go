package controllers

import (
	"context"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"gorm.io/gorm"
)

type PrivacyGrpcController struct {
	service *services.PrivacyService
}

func NewPrivacyController(db *gorm.DB) (*PrivacyGrpcController, error) {
	service, err := services.NewPrivacyService(db)
	if err != nil {
		return nil, err
	}

	return &PrivacyGrpcController{
		service: service,
	}, nil
}

func (s *PrivacyGrpcController) CreatePrivacy(ctx context.Context, in *protopb.CreatePrivacyRequest) (*protopb.EmptyResponsePrivacy, error) {
	var privacy *persistence.Privacy

	privacy.ConvertFromGrpc(in.Privacy)
	_, err := s.service.CreatePrivacy(ctx, privacy)
	if err != nil {
		return &protopb.EmptyResponsePrivacy{}, err
	}

	return &protopb.EmptyResponsePrivacy{}, nil
}

func (s *PrivacyGrpcController) UpdatePrivacy(ctx context.Context, in *protopb.CreatePrivacyRequest) (*protopb.EmptyResponsePrivacy, error) {
	var privacy *persistence.Privacy

	privacy.ConvertFromGrpc(in.Privacy)
	_, err := s.service.UpdatePrivacy(ctx, privacy)
	if err != nil {
		return &protopb.EmptyResponsePrivacy{}, err
	}

	return &protopb.EmptyResponsePrivacy{}, nil
}

func (s *PrivacyGrpcController) BlockUser(ctx context.Context, in *protopb.CreateBlockRequest) (*protopb.EmptyResponsePrivacy, error) {
	var block *persistence.BlockedUsers

	block.ConvertFromGrpc(in.Block)
	_, err := s.service.BlockUser(ctx, block)
	if err != nil {
		return &protopb.EmptyResponsePrivacy{}, err
	}

	return &protopb.EmptyResponsePrivacy{}, nil
}

func (s *PrivacyGrpcController) UnBlockUser(ctx context.Context, in *protopb.CreateBlockRequest) (*protopb.EmptyResponsePrivacy, error) {
	var block *persistence.BlockedUsers

	block.ConvertFromGrpc(in.Block)
	_, err := s.service.UnBlockUser(ctx, block)
	if err != nil {
		return &protopb.EmptyResponsePrivacy{}, err
	}

	return &protopb.EmptyResponsePrivacy{}, nil
}

func (s *PrivacyGrpcController) CheckUserProfilePublic(ctx context.Context, in *protopb.PrivacyRequest) (*protopb.BooleanResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckUserProfilePublic")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	finalResponse := protopb.BooleanResponse{
		Response: s.service.CheckUserProfilePublic(ctx, in.UserId),
	}

	return &finalResponse, nil
}
