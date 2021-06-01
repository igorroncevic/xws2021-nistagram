module github.com/david-drvar/xws2021-nistagram/user_service

go 1.16


replace github.com/david-drvar/xws2021-nistagram/common => ./../common

require (
github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
    github.com/david-drvar/xws2021-nistagram/common v0.0.1
    github.com/golang/protobuf v1.5.2
    github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
    github.com/grpc-ecosystem/grpc-gateway v1.9.5
    github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0 // indirect
    github.com/jackc/pgconn v1.8.1
    github.com/jackc/pgproto3/v2 v2.0.7 // indirect
    github.com/opentracing/opentracing-go v1.2.0
    github.com/uber/jaeger-client-go v2.29.1+incompatible // indirect
    github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
    golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
    golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
    golang.org/x/sys v0.0.0-20210525143221-35b2ab0089ea // indirect
    google.golang.org/genproto v0.0.0-20210524171403-669157292da3
    google.golang.org/grpc v1.38.0
    google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
    google.golang.org/protobuf v1.26.0
    gorm.io/gorm v1.21.10
)
