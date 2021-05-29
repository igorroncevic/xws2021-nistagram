package model

import "time"

type Message struct{
	Id string
	SenderId string
	ReceiverId string
	DateCreated time.Time // TODO
	ContentType string
	IsRead bool
	Content string // interface{}, can be anything
	IsMediaOpened bool
}

type ContentType string
const(
	TypeImage2 ContentType = "Image"
	TypeVideo2             = "Video"
	TypeString             = "String"
	TypeLink               = "Link"
)

type MessageRequest struct{
	SenderId string
	ReceiverId string
	IsAccepted bool
}