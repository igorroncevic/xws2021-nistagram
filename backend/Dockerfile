FROM golang:1.12-alpine as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /build/myapp

# Fetch dependencies
COPY ["user_service", "user_service/"]
COPY ["common" ,"user_service/"]
COPY ["common", "common/"]
WORKDIR /build/myapp/user_service
COPY go.mod ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create final image
FROM alpine
WORKDIR /root
COPY --from=builder /build/myapp/myapp/user_service .
EXPOSE 8001
EXPOSE 8091
CMD ["./myapp"]