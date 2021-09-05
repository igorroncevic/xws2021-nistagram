module github.com/igorroncevic/xws2021-nistagram/monitoring_service

go 1.16

replace github.com/igorroncevic/xws2021-nistagram/monitoring_service => ./monitoring_service

replace github.com/igorroncevic/xws2021-nistagram/common => ./../common

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/igorroncevic/xws2021-nistagram/common v0.0.1
	github.com/jackc/pgconn v1.10.0 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/segmentio/kafka-go v0.4.17 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/uber/jaeger-client-go v2.29.1+incompatible // indirect
	google.golang.org/grpc v1.40.0 // indirect
	gorm.io/driver/postgres v1.1.0 // indirect
	gorm.io/gorm v1.21.14 // indirect
)
