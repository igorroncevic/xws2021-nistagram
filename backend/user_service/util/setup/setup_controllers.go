package setup

import (
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/user_service/controllers"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"gorm.io/gorm"
)

func GetUsersController(db *gorm.DB) controllers.UserController {
	userRepository, _ := repositories.NewUserRepo(db)
	userService := services.UserService{Repository: userRepository}
	userController := controllers.UserController{Service: userService}

	fmt.Println("User controller up and running.")

	return userController
}

func GetPrivacyController(db *gorm.DB) controllers.PrivacyController {
	privacyRepository, _ := repositories.NewPrivacyRepo(db)
	privacyService := services.PrivacyService{Repository: privacyRepository}
	privacyController := controllers.PrivacyController{Service: privacyService}

	fmt.Println("User controller up and running.")

	return privacyController
}