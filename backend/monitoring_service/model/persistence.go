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
