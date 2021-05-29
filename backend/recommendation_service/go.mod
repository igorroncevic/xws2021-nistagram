module github.com/david-drvar/xws2021-nistagram/recommendation_service

go 1.16

replace github.com/david-drvar/xws2021-nistagram/recommendation_service => ./recommendation_service
replace github.com/david-drvar/xws2021-nistagram/common => ./../common

require (
	github.com/david-drvar/xws2021-nistagram/common v0.0.1
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	golang.org/x/text v0.3.6 // indirect
	gorm.io/gorm v1.21.10
)