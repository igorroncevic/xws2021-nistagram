package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
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

func (service *VerificationService) CreateVerificationRequest(ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateVerificationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return "", nil
}
