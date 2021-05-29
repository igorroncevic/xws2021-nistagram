module github.com/david-drvar/xws2021-nistagram/user_service

go 1.16

replace github.com/david-drvar/xws2021-nistagram/user_service => ./user_service
replace github.com/david-drvar/xws2021-nistagram/common => ./../common

require (
	github.com/david-drvar/xws2021-nistagram/common v0.0.1
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgconn v1.8.1
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/text v0.3.6 // indirect
	gorm.io/gorm v1.21.10
)
