package setup

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"xws2021-nistagram/backend/model/persistence"
)

type DbConfig struct {
	DatabaseURL string `json:"database_url"`
}

func InitDatabase(conf DbConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s", conf.DatabaseURL)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	setupDatabase(db)

	fmt.Println("Connected to database.")

	return db
}

func setupDatabase(db *gorm.DB){
	db.AutoMigrate(&persistence.User{},
		&persistence.UserAdditionalInfo{},
		&persistence.Privacy{},
		&persistence.BlockedUsers{},
		&persistence.Followers{},
		&persistence.VerificationRequest{},
		&persistence.APIKeys{},
		)
}