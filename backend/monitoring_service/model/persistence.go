package model

import "time"

type PerformanceMessage struct {
	Id		  string		`gorm:"primaryKey" json:"id"`
	Timestamp time.Time		`json:"timestamp"`
	Service   string		`json:"service"`
	Function  string		`json:"function"`
	Status	  int			`json:"status"`
	Message   string		`json:"message"`
}

type UserEventMessage struct {
	Id		    string		`gorm:"primaryKey"`
	Timestamp   time.Time	`json:"timestamp"`
	Type		string		`json:"type"`
	UserId		string		`json:"user_id"`
	Message		string		`json:"message"`
}
