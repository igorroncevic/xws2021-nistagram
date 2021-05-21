package setup

import (
	"fmt"
	"gorm.io/gorm"
	controllers2 "xws2021-nistagram/backend/controllers"
	repositories2 "xws2021-nistagram/backend/repositories"
	services2 "xws2021-nistagram/backend/services"
)

func GetUsersController(db *gorm.DB) controllers2.UserController {
	userRepository, _ := repositories2.NewUserRepo(db)
	userService := services2.UserService{Repository: userRepository}
	userController := controllers2.UserController{Service: userService}

	fmt.Println("User controller up and running.")

	return userController
}
