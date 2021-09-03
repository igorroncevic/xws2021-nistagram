package main

import (
	"github.com/igorroncevic/xws2021-nistagram/common"
	"github.com/igorroncevic/xws2021-nistagram/common/interceptors/rbac"
	"github.com/igorroncevic/xws2021-nistagram/common/logger"
	"github.com/igorroncevic/xws2021-nistagram/content_service/util/setup"
	"os"
)

func main() {
	if os.Getenv("Docker_env") == "" {
		SetupEnvVariables()
	}

	db := common.InitDatabase(common.ContentDatabase)
	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}

	err = rbac.SetupContentRBAC(db)
	if err != nil {
		panic("Cannot setup rbac tables. Error message: " + err.Error())
	}

	customLogger := logger.NewLogger()
	setup.GRPCServer(db, customLogger)
}

func SetupEnvVariables() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", common.ContentDatabaseName)
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PW", "admin")
}
