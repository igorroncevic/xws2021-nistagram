package persistence

import (
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"

)
func (notification UserNotification) ConvertToGrpc() *protopb.Notification {
	return &protopb.Notification{
		Id:        notification.NotificationId,
		Text:      notification.Text,
		CreatorId: notification.CreatorId,
		UserId:    notification.UserId,
		Type:      notification.Type,
	}
}
