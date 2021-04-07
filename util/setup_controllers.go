package util

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"xws2021-nistagram/controllers"
	"xws2021-nistagram/repositories"
	"xws2021-nistagram/services"
)

type DbConfig struct {
	DatabaseURL string `json:"database_url"`
}

func GetConnectionPool(conf DbConfig) *pgxpool.Pool {
	poolConfig, _ := pgxpool.ParseConfig(conf.DatabaseURL)

	connection, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database.")

	return connection
}

func GetUsersController(connpool *pgxpool.Pool) controllers.UserController {
	userRepository := repositories.NewUserRepo(connpool)
	userService := services.UserService{Repository: userRepository}
	userController := controllers.UserController{Service: userService}

	fmt.Println("User controller up and running.")

	return userController
}
