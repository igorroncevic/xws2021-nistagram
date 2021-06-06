package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
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
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.CreatePrivacy(ctx, privacy)
}

func (service *PrivacyService) UpdatePrivacy(ctx context.Context, privacy *persistence.Privacy) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdatePrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.UpdatePrivacy(ctx, privacy)
}

func (service *PrivacyService) BlockUser(ctx context.Context, block *persistence.BlockedUsers) (bool, error) {
	//TODO Proveri da li ti useri postoje i posalji zahtev da im se obrise prijateljstvo
	span := tracer.StartSpanFromContextMetadata(ctx, "BlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.BlockUser(ctx, block)
}

func (service *PrivacyService) UnBlockUser(ctx context.Context, block *persistence.BlockedUsers) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UnBlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.UnBlockUser(ctx, block)
}

func (service *PrivacyService) CheckUserProfilePublic(ctx context.Context, userId string) bool {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckUserProfilePublic")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var privacy, _ = service.repository.GetUserPrivacy(ctx, userId)
	return privacy.IsProfilePublic == true
}

func (service *PrivacyService) GetAllPublicUsers(ctx context.Context) []string {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPublicUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	privacies, _ := service.repository.GetAlLPublicUsers(ctx)

	publicUsers := []string{}
	for _, privacy := range privacies {
		if privacy.IsProfilePublic {
			publicUsers = append(publicUsers, privacy.UserId)
		}
	}

	return publicUsers
}