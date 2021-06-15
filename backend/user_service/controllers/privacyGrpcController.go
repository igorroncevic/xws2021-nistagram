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
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var privacy *persistence.Privacy
	privacy.ConvertFromGrpc(in.Privacy)
	_, err := s.service.CreatePrivacy(ctx, privacy)
	if err != nil {
		return &protopb.EmptyResponsePrivacy{}, err
	}

	return &protopb.EmptyResponsePrivacy{}, nil
}

func (s *PrivacyGrpcController) UpdatePrivacy(ctx context.Context, in *protopb.CreatePrivacyRequest) (*protopb.EmptyResponsePrivacy, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdatePrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var privacy *persistence.Privacy
	privacy.ConvertFromGrpc(in.Privacy)
	_, err := s.service.UpdatePrivacy(ctx, privacy)
	if err != nil {
		return &protopb.EmptyResponsePrivacy{}, err
	}

	return &protopb.EmptyResponsePrivacy{}, nil
}

func (s *PrivacyGrpcController) BlockUser(ctx context.Context, in *protopb.CreateBlockRequest) (*protopb.EmptyResponsePrivacy, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "BlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var block *persistence.BlockedUsers
	block = block.ConvertFromGrpc(in.Block)
	_, err := s.service.BlockUser(ctx, block)
	if err != nil {
		return &protopb.EmptyResponsePrivacy{}, err
	}

	return &protopb.EmptyResponsePrivacy{}, nil
}

func (s *PrivacyGrpcController) UnBlockUser(ctx context.Context, in *protopb.CreateBlockRequest) (*protopb.EmptyResponsePrivacy, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UnBlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var block *persistence.BlockedUsers
	block = block.ConvertFromGrpc(in.Block)
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

func (s *PrivacyGrpcController) CheckIfBlocked(ctx context.Context, in *protopb.CreateBlockRequest) (*protopb.BooleanResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckIfBlocked")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	isBlocked, err := s.service.CheckIfBlocked(ctx, in.Block.UserId, in.Block.BlockedUserId)
	if err != nil { return &protopb.BooleanResponse{Response: true}, nil }

	return &protopb.BooleanResponse{
		Response: isBlocked,
	}, nil
}

func (s *PrivacyGrpcController) GetAllPublicUsers(ctx context.Context, in *protopb.RequestIdPrivacy) (*protopb.StringArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPublicUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	users := s.service.GetAllPublicUsers(ctx)
	if in.Id == "" || in == nil {
		return &protopb.StringArray{
			Ids: users,
		}, nil
	}

	nonBlockedUsers := []string{}
	for _, user := range users {
		isBlocked, err := s.service.CheckIfBlocked(ctx, user, in.Id)
		if err == nil || !isBlocked {
			nonBlockedUsers = append(nonBlockedUsers, user)
		}
	}

	return &protopb.StringArray{
		Ids: nonBlockedUsers,
	}, nil
}
