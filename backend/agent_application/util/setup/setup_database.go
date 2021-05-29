package setup

import (
	"github.com/david-drvar/xws2021-nistagram/agent_application/model"
	"gorm.io/gorm"
)

func FillDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{},
		&model.Product{},
		&model.Order{},
		&model.OrderProducts{},
		&model.PostReport{},
		&model.StoryReport{},
	)

	return err
}
