module github.com/david-drvar/xws2021-nistagram/chat_service

go 1.16

replace github.com/david-drvar/xws2021-nistagram/common => ./../common

require (
	github.com/david-drvar/xws2021-nistagram/common v0.0.1
	github.com/gin-gonic/gin v1.7.2
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gorm.io/gorm v1.21.9

)
