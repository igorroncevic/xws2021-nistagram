package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"gorm.io/gorm"
)

type PrivacyRepository interface {
	CreatePrivacy(ctx context.Context, privacy *persistence.Privacy) (persistence.Privacy, error)
	UpdatePrivacy(ctx context.Context, privacy *persistence.Privacy) (bool, error)
	BlockUser(ctx context.Context, block *persistence.BlockedUsers) (bool, error)
	UnBlockUser(ctx context.Context, block *persistence.BlockedUsers) (bool, error)
}

type privacyRepository struct {
	DB *gorm.DB
}

func NewPrivacyRepo(db *gorm.DB) (PrivacyRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}

	return &privacyRepository{DB: db}, nil
}

func (repository *privacyRepository) BlockUser(ctx context.Context, b *persistence.BlockedUsers) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "BlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	db := repository.DB.Create(&b)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *privacyRepository) UnBlockUser(ctx context.Context, b *persistence.BlockedUsers) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UnBlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	db := repository.DB.Delete(&b)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *privacyRepository) CreatePrivacy(ctx context.Context, p *persistence.Privacy) (persistence.Privacy, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Create(&p)
	return *p, err.Error
}

func (repository *privacyRepository) UpdatePrivacy(ctx context.Context, p *persistence.Privacy) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdatePrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var privacy persistence.Privacy

	db := repository.DB.Model(&privacy).Where("user_id = ?", p.UserId).Updates(persistence.Privacy{IsTagEnabled: p.IsTagEnabled, IsProfilePublic: p.IsProfilePublic, IsDMPublic: p.IsDMPublic})
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}
