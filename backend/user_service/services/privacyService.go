package services

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
)

type PrivacyService struct {
	Repository repositories.PrivacyRepository
}

func (service *PrivacyService) CreatePrivacy(privacy *persistence.Privacy) (persistence.Privacy, error) {
	return service.Repository.CreatePrivacy(privacy);
}
