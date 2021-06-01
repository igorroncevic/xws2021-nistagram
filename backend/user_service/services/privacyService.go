package services

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"gorm.io/gorm"
)

type PrivacyService struct {
	repository repositories.PrivacyRepository
}

func NewPrivacyService(db *gorm.DB) (*PrivacyService, error){
	repository, err := repositories.NewPrivacyRepo(db)

	return &PrivacyService{
		repository: repository,
	}, err
}

func (service *PrivacyService) CreatePrivacy(privacy *persistence.Privacy) (persistence.Privacy, error) {
	return service.repository.CreatePrivacy(privacy)
}

func (service *PrivacyService) UpdatePrivacy(privacy *persistence.Privacy) (bool, error) {
	return service.repository.UpdatePrivacy(privacy)
}
