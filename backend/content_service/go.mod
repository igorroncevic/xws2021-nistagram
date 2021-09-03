module github.com/igorroncevic/xws2021-nistagram/content_service

go 1.16

replace github.com/igorroncevic/xws2021-nistagram/content_service => ./content_service

replace github.com/igorroncevic/xws2021-nistagram/common => ./../common

require (
	github.com/dvwright/xss-mw v0.0.0-20191029162136-7a0dab86d8f6 // indirect
	github.com/gin-gonic/gin v1.7.2 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/igorroncevic/xws2021-nistagram/common v0.0.1
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/microcosm-cc/bluemonday v1.0.9 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/uber/jaeger-client-go v2.29.1+incompatible // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20210831024726-fe130286e0e2
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.21.14
)
