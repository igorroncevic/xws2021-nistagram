package repositories

import (
	"errors"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"gorm.io/gorm"
)

type PrivacyRepository interface {
	CreatePrivacy(privacy *persistence.Privacy) (persistence.Privacy, error)
	UpdatePrivacy(privacy *persistence.Privacy) (bool, error)
}

type privacyRepository struct {
	DB *gorm.DB
}

func NewPrivacyRepo(db *gorm.DB) (PrivacyRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}

	return &privacyRepository{ DB: db }, nil
}

func (repository *privacyRepository) CreatePrivacy(p *persistence.Privacy) (persistence.Privacy, error) {

	return *p ,nil
}

func (repository *privacyRepository) UpdatePrivacy(p *persistence.Privacy)  (bool, error){
	var privacy persistence.Privacy

	db := repository.DB.Model(&privacy).Where("user_id = ?", p.UserId).Updates(persistence.Privacy{IsTagEnabled: p.IsTagEnabled, IsProfilePublic: p.IsProfilePublic, IsDMPublic: p.IsDMPublic})
	if db.Error != nil {
		return false, db.Error
	}else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}