package setup

import (
	"gorm.io/gorm"
	"xws2021-nistagram/backend/model/persistence"
)

func SetupDatabase(db *gorm.DB){
	db.AutoMigrate(&persistence.User{},
		&persistence.UserAdditionalInfo{},
		&persistence.Privacy{},
		&persistence.BlockedUsers{},
		&persistence.Followers{},
		&persistence.VerificationRequest{},
		&persistence.APIKeys{},
		)
}