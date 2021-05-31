package repositories

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"gorm.io/gorm"
)

type PrivacyRepository interface {
	CreatePrivacy(privacy *persistence.Privacy) (persistence.Privacy, error)
	UpdateUserPrivacy(privacy persistence.Privacy) (bool, error)
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

func (repository *privacyRepository) CreatePrivacy(privacy *persistence.Privacy) (persistence.Privacy, error) {

	return *privacy ,nil
}

func (repository *privacyRepository) UpdateUserPrivacy(privacy persistence.Privacy)  (bool, error){

	return false, nil
}