package main

import (
	"github.com/david-drvar/xws2021-nistagram/agent_application/util/setup"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/interceptors/rbac"
	"os"
)

func main() {
	if os.Getenv("Docker_env") == "" {
		SetupEnvVariables()
	}

	db := common.InitDatabase(common.AgentDatabase)

	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}

	err = rbac.SetupAgentsRBAC(db)
	if err != nil {
		panic("Cannot setup rbac tables. Error message: " + err.Error())
	}

	setup.GRPCServer(db)
}

func SetupEnvVariables() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", common.AgentDatabaseName)
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PW", "root")
}
