package kafka_util

import "time"

const (
	ExampleGroupId         = "groupId"
	PerformanceTopic       = "performance"
	UserEventsTopic        = "user-events"
	RetryTopic             = "retry"
	RegularConsumerMaxWait = time.Duration(10) * time.Second
	RetryConsumerMaxWait   = time.Duration(5) * time.Second

	/* Services */
	UserService    = "UserService"
	ContentService = "ContentService"

	/* Functions */
	LoginFunction              = "Login"
	CreateNotificationFunction = "CreateNotification"
	GenerateApiTokenFunction   = "GenerateApiToken"
)
