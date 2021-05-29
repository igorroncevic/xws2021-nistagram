package common

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
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

	fmt.Println("Connected to database.")

	return db
}


