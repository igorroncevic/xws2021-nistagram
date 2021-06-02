module github.com/david-drvar/xws2021-nistagram/recommendation_service

go 1.16

replace github.com/david-drvar/xws2021-nistagram/recommendation_service => ./recommendation_service

replace github.com/david-drvar/xws2021-nistagram/common => ./../common

require (
	github.com/david-drvar/xws2021-nistagram/common v0.0.1
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/text v0.3.6 // indirect
    github.com/neo4j/neo4j-go-driver/v4 v4.3.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/gorm v1.21.10
)
