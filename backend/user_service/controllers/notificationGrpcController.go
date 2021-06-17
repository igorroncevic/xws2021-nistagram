package controllers

import (
	"context"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"gorm.io/gorm"
)

type NotificationGrpcController struct {
	service *services.NotificationService
}

func NewNotificationController(db *gorm.DB) (*NotificationGrpcController, error) {
	service, err := services.NewNotificationService(db)
	if err != nil {
		return nil, err
	}

	return &NotificationGrpcController{
		service: service,
	}, nil
}

func (c *NotificationGrpcController) CreateNotification(ctx context.Context, in *protopb.CreateNotificationRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var notification *domain.UserNotification
	notification = notification.ConvertFromGrpc(*in)
	err := c.service.CreateNotification(ctx, notification)
	if err != nil {
		return &protopb.EmptyResponse{}, err
	}

	return &protopb.EmptyResponse{}, nil
}

func (c *NotificationGrpcController) GetUserNotifications(ctx context.Context, in *protopb.RequestIdUsers) (*protopb.CreateNotificationResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserNotifications")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	notifications, err := c.service.GetUserNotifications(ctx, in.Id)
	if err != nil {
		return &protopb.CreateNotificationResponse{}, err
	}

	var response []*protopb.Notification
	for _, n := range notifications {
		response = append(response, n.ConvertToGrpc())
	}

	return &protopb.CreateNotificationResponse{
		Notifications: response,
	}, nil
}
