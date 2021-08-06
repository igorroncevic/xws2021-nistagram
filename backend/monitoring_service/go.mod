module github.com/igorroncevic/xws2021-nistagram/monitoring_service

go 1.16

replace github.com/igorroncevic/xws2021-nistagram/monitoring_service => ./monitoring_service

replace github.com/igorroncevic/xws2021-nistagram/common => ./../common

require (
	github.com/igorroncevic/xws2021-nistagram/common v0.0.1
	github.com/segmentio/kafka-go v0.4.17 // indirect
)
