module github.com/igorroncevic/xws2021-nistagram/content_service

go 1.16

replace github.com/igorroncevic/xws2021-nistagram/content_service => ./content_service

replace github.com/igorroncevic/xws2021-nistagram/common => ./../common

require (
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/igorroncevic/xws2021-nistagram/common v0.0.1
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.11.0
	github.com/satori/go.uuid v1.2.0
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.21.14
)
