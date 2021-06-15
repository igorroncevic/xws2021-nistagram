package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/util"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/images"
	"gorm.io/gorm"
	"time"
)

type VerificationRepository interface {
	CreateVerificationRequest(context.Context, domain.VerificationRequest) error
	SaveUserDocumentPhoto(ctx context.Context, verificationRequest domain.VerificationRequest) (string, error)
	GetPendingVerificationRequests(ctx context.Context) ([]domain.VerificationRequest, error)
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

func (repository *verificationRepository) CreateVerificationRequest(ctx context.Context, verificationRequest domain.VerificationRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateVerificationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		var userAdditionalInfo persistence.UserAdditionalInfo

		result := repository.DB.Model(&userAdditionalInfo).Where("id = ?", verificationRequest.UserId).Updates(persistence.UserAdditionalInfo{
			Category: verificationRequest.Category,
		})
		if result.Error != nil || result.RowsAffected != 1 {
			return errors.New("cannot update user additional info")
		}

		var documentPhotoDecoded, resultImage = repository.SaveUserDocumentPhoto(ctx, verificationRequest)
		if resultImage != nil {
			return errors.New("cannot decode document photo")
		}

		var verificationRequestPersistence = persistence.VerificationRequest{
			UserId:        verificationRequest.UserId,
			DocumentPhoto: documentPhotoDecoded,
			Status:        model.Pending,
			CreatedAt:     time.Time{},
		}
		result = repository.DB.Create(&verificationRequestPersistence)
		if result.Error != nil {
			return errors.New("cannot create verification request")
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (repository *verificationRepository) SaveUserDocumentPhoto(ctx context.Context, verificationRequest domain.VerificationRequest) (string, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveUserDocumentPhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	mimeType, err := images.GetImageType(verificationRequest.DocumentPhoto)
	if err != nil {
		return "", err
	}

	t := time.Now()
	formatted := fmt.Sprintf("%s%d%02d%02d%02d%02d%02d%02d", verificationRequest.UserId, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	name := formatted + "." + mimeType

	err = images.SaveImage(name, verificationRequest.DocumentPhoto)
	if err != nil {
		return "", err
	}

	verificationRequest.DocumentPhoto = util.GetContentLocation(name)
	return verificationRequest.DocumentPhoto, nil
}

func (repository *verificationRepository) GetPendingVerificationRequests(ctx context.Context) ([]domain.VerificationRequest, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPendingVerificationRequests")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var verificationRequestsPersistence []persistence.VerificationRequest
	result := repository.DB.Where("status = ?", model.Pending).Find(&verificationRequestsPersistence)
	if result.Error != nil {
		return nil, result.Error
	}

	var verificationRequestsDomain []domain.VerificationRequest

	for _, verificationRequest := range verificationRequestsPersistence {
		user, err := repository.userRepository.GetUserById(ctx, verificationRequest.UserId)
		if err != nil {
			return nil, err
		}
		userAdditionalInfo, err := repository.userRepository.GetUserAdditionalInfoById(ctx, verificationRequest.UserId)
		if err != nil {
			return nil, err
		}
		imageBase64, err := images.LoadImageToBase64(verificationRequest.DocumentPhoto)
		if err != nil {
			return nil, err
		}
		verificationRequestsDomain = append(verificationRequestsDomain, domain.VerificationRequest{
			UserId:        verificationRequest.UserId,
			DocumentPhoto: imageBase64,
			Status:        verificationRequest.Status,
			CreatedAt:     verificationRequest.CreatedAt,
			Category:      userAdditionalInfo.Category,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
		})
	}

	return verificationRequestsDomain, nil
}
