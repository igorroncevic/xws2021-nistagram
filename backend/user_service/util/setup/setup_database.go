package setup

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"gorm.io/gorm"
)

func FillDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{},
		&model.UserAdditionalInfo{},
		&model.Privacy{},
		&model.BlockedUsers{},
		&model.Followers{},
		&model.VerificationRequest{},
		&model.APIKeys{},
	)

	return err
}