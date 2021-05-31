package main

import (
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/setup"
)

func main() {
	db := common.InitDatabase()
	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}

	setup.GRPCServer(db)
}