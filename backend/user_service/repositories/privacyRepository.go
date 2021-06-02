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
	UpdatePrivacy(privacy *persistence.Privacy) (bool, error)
	BlockUser(block *persistence.BlockedUsers) (bool, error)
	UnBlockUser(block *persistence.BlockedUsers) (bool, error)
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

func (repository *privacyRepository) BlockUser(b *persistence.BlockedUsers) (bool, error) {
	db := repository.DB.Create(&b)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *privacyRepository) UnBlockUser(b *persistence.BlockedUsers) (bool, error) {
	db := repository.DB.Delete(&b)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}

func (repository *privacyRepository) CreatePrivacy(ctx context.Context, p *persistence.Privacy) (persistence.Privacy, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Create(&p)
	return *p, err.Error
}

func (repository *privacyRepository) UpdatePrivacy(p *persistence.Privacy) (bool, error) {
	var privacy persistence.Privacy

	db := repository.DB.Model(&privacy).Where("user_id = ?", p.UserId).Updates(persistence.Privacy{IsTagEnabled: p.IsTagEnabled, IsProfilePublic: p.IsProfilePublic, IsDMPublic: p.IsDMPublic})
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}
