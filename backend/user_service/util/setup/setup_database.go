package setup

import (
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"gorm.io/gorm"
)

func FillDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&persistence.User{},
		&persistence.UserAdditionalInfo{},
		&persistence.Privacy{},
		&persistence.BlockedUsers{},
		&persistence.Followers{},
		&persistence.VerificationRequest{},
		&persistence.APIKeys{},
		&persistence.UserNotification{},
		)

	return err
}