package util

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	controllers2 "xws2021-nistagram/backend/controllers"
	repositories2 "xws2021-nistagram/backend/repositories"
	services2 "xws2021-nistagram/backend/services"
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

func GetUsersController(connpool *pgxpool.Pool) controllers2.UserController {
	userRepository := repositories2.NewUserRepo(connpool)
	userService := services2.UserService{Repository: userRepository}
	userController := controllers2.UserController{Service: userService}

	fmt.Println("User controller up and running.")

	return userController
}
