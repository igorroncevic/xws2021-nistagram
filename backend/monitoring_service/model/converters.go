package model

import (
	"github.com/igorroncevic/xws2021-nistagram/common/kafka_util"
	"time"
)

func ConvertPerformanceMessageToPersistence(id string, timestamp time.Time, message kafka_util.PerformanceMessage) PerformanceMessage {
	return PerformanceMessage{
		Id:        id,
		Timestamp: timestamp,
		Service:   message.Service,
		Function:  message.Function,
		Status:    message.Status,
		Message:   message.Message,
	}
}

func ConvertUserEventMessageToPersistence(id string, timestamp time.Time, message kafka_util.UserEventMessage) UserEventMessage {
	return UserEventMessage{
		Id:        id,
		Timestamp: timestamp,
		Type:      message.Type.String(),
		UserId:    message.UserId,
		Message:   message.Message,
	}
}
