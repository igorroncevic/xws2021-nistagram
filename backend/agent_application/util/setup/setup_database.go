package setup

import (
	"github.com/igorroncevic/xws2021-nistagram/agent_application/model/persistence"
	"gorm.io/gorm"
)

func FillDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&persistence.User{},
		&persistence.Product{},
		&persistence.Order{},
		&persistence.OrderProducts{},
		&persistence.PostReport{},
		&persistence.StoryReport{},
		&persistence.APIKey{},
	)

	return err
}
