package persistence

import (
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"

)
func (n UserNotification) ConvertToGrpc() *protopb.Notification {
	return &protopb.Notification{
		Id: n.NotificationId,
		Text: n.Text,
		CreatorId: n.CreatorId,
		UserId: n.UserId,
		Type: n.Type,
	}
}
