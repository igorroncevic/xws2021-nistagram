package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"gorm.io/gorm"
)

type VerificationService struct {
	userRepository         repositories.UserRepository
	verificationRepository repositories.VerificationRepository
}

func NewVerificationService(db *gorm.DB) (*VerificationService, error) {
	userRepository, err := repositories.NewUserRepo(db)
	verificationRepository, err := repositories.NewVerificationRepo(db)

	return &VerificationService{
		userRepository,
		verificationRepository,
	}, err
}

func (service *VerificationService) CreateVerificationRequest(ctx context.Context, verificationRequest domain.VerificationRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateVerificationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.verificationRepository.CreateVerificationRequest(ctx, verificationRequest)
	if err != nil {
		return err
	}

	return nil
}

func (service *VerificationService) GetPendingVerificationRequests(ctx context.Context) ([]domain.VerificationRequest, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPendingVerificationRequests")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	verificationRequests, err := service.verificationRepository.GetPendingVerificationRequests(ctx)
	if err != nil {
		return nil, err
	}

	return verificationRequests, nil
}
