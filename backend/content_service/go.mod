module github.com/david-drvar/xws2021-nistagram/content_service

go 1.16

replace github.com/david-drvar/xws2021-nistagram/content_service => ./content_service

replace github.com/david-drvar/xws2021-nistagram/common => ./../common

require (
	github.com/david-drvar/xws2021-nistagram/common v0.0.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/satori/go.uuid v1.2.0
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20210426193834-eac7f76ac494
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gorm.io/gorm v1.21.10
)
