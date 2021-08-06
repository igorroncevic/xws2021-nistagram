package main

import (
	"github.com/igorroncevic/xws2021-nistagram/common"
	"github.com/igorroncevic/xws2021-nistagram/common/interceptors/rbac"
	"github.com/igorroncevic/xws2021-nistagram/user_service/util/setup"
	"os"
)

func main() {
	if os.Getenv("Docker_env") == "" {
		SetupEnvVariables()
	}

	db := common.InitDatabase(common.UserDatabase)

	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}

	err = rbac.SetupUsersRBAC(db)
	if err != nil {
		panic("Cannot setup rbac tables. Error message: " + err.Error())
	}

	setup.GRPCServer(db)
}

func SetupEnvVariables() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", common.UsersDatabaseName)
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PW", "root")
}
