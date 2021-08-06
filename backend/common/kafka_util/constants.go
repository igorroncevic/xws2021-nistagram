package kafka_util

import "time"

const (
	exampleTopic        	= "exampleTopic"
	RetryTopic		        = "retry"
	RegularConsumerMaxWait  = time.Duration(10) * time.Second
	RetryConsumerMaxWait	= time.Duration(5)  * time.Second
)