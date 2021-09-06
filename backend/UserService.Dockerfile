FROM golang:1.16.3 as builder
WORKDIR /build/myapp

# Fetch dependencies
COPY ["user_service", "github.com/igorroncevic/xws2021-nistagram/user_service/"]
COPY ["common", "github.com/igorroncevic/xws2021-nistagram/common/"]
WORKDIR /build/myapp/github.com/igorroncevic/xws2021-nistagram/user_service
RUN GOPROXY=https://proxy.golang.org/ GO111MODULE=on go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create final image
FROM alpine
RUN apk --no-cache add ca-certificates git
WORKDIR /root
COPY --from=builder /build/myapp/github.com/igorroncevic/xws2021-nistagram/user_service .
EXPOSE 8001
EXPOSE 8091
CMD ["./main"]