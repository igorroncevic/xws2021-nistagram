package common

import (
	"fmt"
	"github.com/lytics/confl"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type dbConfig struct {
	DatabaseURL string `json:"database_url"`
}

func InitDatabase() *gorm.DB {
	var dbConf dbConfig
	if _, err := confl.DecodeFile("./../dbconfig.conf", &dbConf); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s", dbConf.DatabaseURL)
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


