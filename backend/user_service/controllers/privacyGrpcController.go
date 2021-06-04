package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	userspb "github.com/david-drvar/xws2021-nistagram/user_service/proto"
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

func (s *PrivacyGrpcController) CreatePrivacy(ctx context.Context, in *userspb.CreatePrivacyRequest) (*userspb.EmptyResponsePrivacy, error) {
	var privacy *persistence.Privacy

	privacy.ConvertFromGrpc(in.Privacy)
	_, err := s.service.CreatePrivacy(ctx, privacy)
	if err != nil {
		return &userspb.EmptyResponsePrivacy{}, err
	}

	return &userspb.EmptyResponsePrivacy{}, nil
}

func (s *PrivacyGrpcController) UpdatePrivacy(ctx context.Context, in *userspb.CreatePrivacyRequest) (*userspb.EmptyResponsePrivacy, error) {
	var privacy *persistence.Privacy

	privacy.ConvertFromGrpc(in.Privacy)
	_, err := s.service.UpdatePrivacy(ctx, privacy)
	if err != nil {
		return &userspb.EmptyResponsePrivacy{}, err
	}

	return &userspb.EmptyResponsePrivacy{}, nil
}

func (s *PrivacyGrpcController) BlockUser(ctx context.Context, in *userspb.CreateBlockRequest) (*userspb.EmptyResponsePrivacy, error) {
	var block *persistence.BlockedUsers

	block.ConvertFromGrpc(in.Block)
	_, err := s.service.BlockUser(ctx, block)
	if err != nil {
		return &userspb.EmptyResponsePrivacy{}, err
	}

	return &userspb.EmptyResponsePrivacy{}, nil
}

func (s *PrivacyGrpcController) UnBlockUser(ctx context.Context, in *userspb.CreateBlockRequest) (*userspb.EmptyResponsePrivacy, error) {
	var block *persistence.BlockedUsers

	block.ConvertFromGrpc(in.Block)
	_, err := s.service.UnBlockUser(ctx, block)
	if err != nil {
		return &userspb.EmptyResponsePrivacy{}, err
	}

	return &userspb.EmptyResponsePrivacy{}, nil
}
