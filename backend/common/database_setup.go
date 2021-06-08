package common

import (
	"fmt"
	"github.com/lytics/confl"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

type dbConfig struct {
	UserDatabaseURL string `json:"user_database_url"`
	ContentDatabaseURL string `json:"content_database_url"`
	AgentDatabaseURL string `json:"agent_database_url"`
	RecommendationDatabaseURL string `json:"recommendation_database_url"`
}

const (
	UserDatabase 		   = "UserDatabase"
	ContentDatabase 	   = "ContentDatabase"
	AgentDatabase 		   = "AgentDatabase"
	RecommendationDatabase = "RecommendationDatabase"
)

func InitDatabase(dbname string) *gorm.DB {
	var dbConf dbConfig
	if _, err := confl.DecodeFile("./../dbconfig.conf", &dbConf); err != nil {
		panic(err)
	}

	var dsn string
	if dbname == UserDatabase {
		dsn = fmt.Sprintf("%s", dbConf.UserDatabaseURL)
	}else if dbname == ContentDatabase {
		dsn = fmt.Sprintf("%s", dbConf.ContentDatabaseURL)
	}else if dbname == AgentDatabase {
		dsn = fmt.Sprintf("%s", dbConf.AgentDatabaseURL)
	}else if dbname == RecommendationDatabase {
		dsn = fmt.Sprintf("%s", dbConf.RecommendationDatabaseURL)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database.")

	return db
}


