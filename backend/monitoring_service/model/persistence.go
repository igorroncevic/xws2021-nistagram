package model

import "time"

type PerformanceMessage struct {
	Id		  string		`gorm:"primaryKey"`
	Timestamp time.Time
	Service   string
	Function  string
	Status	  int
	Message   string
}

type UserEventMessage struct {
	Id		    string		`gorm:"primaryKey"`
	Timestamp   time.Time
	Type		string
	UserId		string
	Message		string
}
