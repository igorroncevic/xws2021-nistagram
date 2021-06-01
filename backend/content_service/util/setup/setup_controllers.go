package setup

import (
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/content_service/controllers"
	"gorm.io/gorm"
)

func GetContentController(db *gorm.DB) *controllers.ContentGrpcController {
	contentController, _ := controllers.NewContentController(db)

	fmt.Println("Content controller up and running.")

	return contentController
}