package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"gorm.io/gorm"
)

type NotificationService struct {
	repository  repositories.NotificationRepository
	userService *UserService
}

func NewNotificationService(db *gorm.DB) (*NotificationService, error) {
	repository, err := repositories.NewNotificationRepo(db)
	service, err := NewUserService(db)
	return &NotificationService{
		repository:  repository,
		userService: service,
	}, err
}


func (s NotificationService) CreateNotification(ctx context.Context, domainNotification *domain.UserNotification) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	userNotification := &persistence.UserNotification{}

	if domainNotification.NotificationType == "Message" {
		userNotification.Text = domainNotification.CreatorId + " send you a message."
	}else if domainNotification.NotificationType == "Follow" {
		userNotification.Text = domainNotification.CreatorId + " started following you."
	}else if domainNotification.NotificationType == "Like" {
		userNotification.Text = domainNotification.CreatorId + " liked your post."
	}else if domainNotification.NotificationType == "Dislike" {
		userNotification.Text = domainNotification.CreatorId + " disliked your post."
	}else if domainNotification.NotificationType == "Comment" {
		userNotification.Text = domainNotification.CreatorId + " commented on your post."
	}else {
		return errors.New("Bad notification type")
	}
	userNotification.UserId = domainNotification.UserId
	userNotification.CreatorId = domainNotification.CreatorId
	userNotification.Type = domainNotification.NotificationType

	err := s.repository.CreateNotification(ctx, userNotification)
	if err != nil {
		return err
	}

	return nil
}
