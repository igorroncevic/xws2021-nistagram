package kafka_util

import (
	"encoding/json"
	"errors"
)

type PerformanceMessage struct {
	Service   string		`json:"service"`
	Function  string		`json:"function"`
	Status	  int			`json:"status"`
	Message   string		`json:"message"`
}

func NewPerformanceMessage(service string, function string, status int, message string) PerformanceMessage{
	return PerformanceMessage{
		Service:   service,
		Function:  function,
		Status:    status,
		Message:   message,
	}
}

func MarshalPerformanceMessage(service string, function string, status int, message string) (string, error) {
	performanceMessage := NewPerformanceMessage(service, function, status, message)

	jsonMessage, err := json.Marshal(performanceMessage)
	if err != nil {
		return "", errors.New("could not marshal performance message")
	}

	return string(jsonMessage), nil
}

type UserEventsMessage struct {
	Type		UserEventType	`json:"type"`
	UserId		string			`json:"userId"`
	Message		string			`json:"message"`
}

func NewUserEventsMessage(eventType UserEventType, userId string, message string) UserEventsMessage{
	return UserEventsMessage{
		Type:      eventType,
		UserId:    userId,
		Message:   message,
	}
}

type UserEventType string
const (
	LinkClick UserEventType = "LinkClick"
	TimeSpent				= "TimeSpent"
	Login					= "Login"
)


