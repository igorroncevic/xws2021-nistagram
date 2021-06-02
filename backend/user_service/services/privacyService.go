package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"gorm.io/gorm"
)

type PrivacyService struct {
	repository  repositories.PrivacyRepository
	userService *UserService
}

func NewPrivacyService(db *gorm.DB) (*PrivacyService, error) {
	repository, err := repositories.NewPrivacyRepo(db)
	service, err := NewUserService(db)
	return &PrivacyService{
		repository:  repository,
		userService: service,
	}, err
}

func (service *PrivacyService) CreatePrivacy(ctx context.Context, privacy *persistence.Privacy) (persistence.Privacy, error) {
	return service.repository.CreatePrivacy(ctx, privacy)
}

func (service *PrivacyService) UpdatePrivacy(privacy *persistence.Privacy) (bool, error) {
	return service.repository.UpdatePrivacy(privacy)
}

func (service *PrivacyService) BlockUser(block *persistence.BlockedUsers) (bool, error) {
	//TODO Proveri da li ti useri postoje i posalji zahtev da im se obrise prijateljstvo
	return service.repository.BlockUser(block)
}

func (service *PrivacyService) UnBlockUser(block *persistence.BlockedUsers) (bool, error) {
	return service.repository.UnBlockUser(block)
}
