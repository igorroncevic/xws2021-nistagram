package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	CreateNotification(context.Context, *persistence.UserNotification) error
	GetUserNotifications( context.Context,  string) ([]persistence.UserNotification, error)
	DeleteNotification(ctx context.Context,   id string ) (bool, error)
}
type notificationRepository struct {
	DB *gorm.DB
}

func NewNotificationRepo(db *gorm.DB) (NotificationRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}

	return &notificationRepository{DB: db}, nil
}

func (repository *notificationRepository) CreateNotification(ctx context.Context, notification *persistence.UserNotification) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	notification.NotificationId = uuid.New().String()
	result := repository.DB.Create(&notification)
	if result.Error != nil || result.RowsAffected != 1 {
		return errors.New("Could not create notification in repository")
	}
	return nil
}
func (repository *notificationRepository) GetUserNotifications(ctx context.Context,userId string) ([]persistence.UserNotification, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserNotifications")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var userNotifications []persistence.UserNotification

	result := repository.DB.Where("user_id = ?", userId).Find(&userNotifications)
	if result.Error != nil {
		return nil, errors.New("Error while loading notifications")
	}

	return userNotifications, nil
}

func (repository *notificationRepository) GetNotificationById(ctx context.Context, id string) (persistence.UserNotification, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var notification persistence.UserNotification
	result := repository.DB.Where("notification_id = ?", id).Find(&notification)
	if result.Error != nil || result.RowsAffected != 1 {
		return persistence.UserNotification{}, errors.New("cannot retrieve this notification")
	}

	return notification, nil
}
func (repository *notificationRepository) DeleteNotification(ctx context.Context, id string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UnBlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var notification persistence.UserNotification
	notification, _ = repository.GetNotificationById(ctx, id)

	db := repository.DB.Delete(&notification)
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected == 0 {
		return false, errors.New("rows affected is equal to zero")
	}

	return true, nil
}
