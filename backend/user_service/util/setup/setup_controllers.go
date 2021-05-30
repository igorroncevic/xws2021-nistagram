package setup

import (
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/user_service/controllers"
	"gorm.io/gorm"
)

func GetUsersController(db *gorm.DB) *controllers.UserGrpcController {
	userController, _ := controllers.NewUserController(db)

	fmt.Println("User controller up and running.")

	return userController
}
