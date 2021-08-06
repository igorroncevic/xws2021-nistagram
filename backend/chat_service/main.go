package main

import (
	"github.com/igorroncevic/xws2021-nistagram/chat_service/controllers"
	"github.com/igorroncevic/xws2021-nistagram/chat_service/util/setup"
	"github.com/igorroncevic/xws2021-nistagram/common"
	"os"
)

func main() {
	if os.Getenv("Docker_env") == "" {
		SetupEnvVariables()
	}
	db := common.InitDatabase(common.ChatDatabase)
	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}
	controller, _ := controllers.NewMessageController(db)
	setup.ServerSetup(controller)
}

func SetupEnvVariables() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", common.ChatDatabaseName)
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PW", "root")
}
