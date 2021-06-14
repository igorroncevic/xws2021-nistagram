package repositories

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"gorm.io/gorm"
)

type VerificationRepository interface {
	CreateVerificationRequest(context.Context) (bool, error)
}

type verificationRepository struct {
	DB             *gorm.DB
	userRepository UserRepository
}

func NewVerificationRepo(db *gorm.DB) (*verificationRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}
	userRepository, _ := NewUserRepo(db)

	return &verificationRepository{DB: db, userRepository: userRepository}, nil
}

func (repository *verificationRepository) CreateVerificationRequest(ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateVerificationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return true, nil
}
