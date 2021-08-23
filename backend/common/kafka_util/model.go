package kafka_util

import (
	"encoding/json"
	"errors"
	"fmt"
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

type UserEventMessage struct {
	Type		UserEventType	`json:"type"`
	UserId		string			`json:"userId"`
	Message		string			`json:"message"`
}

func NewUserEventsMessage(eventType UserEventType, userId string, message string) UserEventMessage {
	return UserEventMessage{
		Type:      eventType,
		UserId:    userId,
		Message:   message,
	}
}

func MarshalUserEventMessage(eventType UserEventType, userId string, message string) (string, error) {
	performanceMessage := NewUserEventsMessage(eventType, userId, message)

	jsonMessage, err := json.Marshal(performanceMessage)
	if err != nil {
		return "", errors.New("could not marshal user event message")
	}

	return string(jsonMessage), nil
}

type UserEventType string
const (
	LinkClick UserEventType = "LinkClick"
	TimeSpent				= "TimeSpent"
	Login					= "Login"
)

func (uet UserEventType) String() string {
	switch uet {
	case LinkClick:
		return "LinkClick"
	case TimeSpent:
		return "TimeSpent"
	case Login:
		return "Login"
	default:
		return fmt.Sprintf("%s", string(uet))
	}
}

func GetPostType(uet string) UserEventType {
	switch uet {
	case "LinkClick", "linkClick":
		return LinkClick
	case "TimeSpent", "timeSpent":
		return TimeSpent
	case "Login", "login":
		return Login
	default:
		return ""
	}
}

