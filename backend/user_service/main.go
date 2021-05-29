package main

import (
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/user_service/util"
	"github.com/gorilla/mux"
	"github.com/lytics/confl"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	var dbConf common.DbConfig
	if _, err := confl.DecodeFile("./../dbconfig.conf", &dbConf); err != nil {
		panic(err)
	}

	db := common.InitDatabase(dbConf)
	userController := util.GetUsersController(db)

	r := mux.NewRouter()

	usersRouter := r.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", userController.GetAllUsers).Methods("GET")

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", userController.CreateUser).Methods("POST")
	authRouter.HandleFunc("/login", userController.LoginUser).Methods("POST")

	usersRouter.Use(common.AuthMiddleware) // Authenticate user's JWT before letting them access these endpoints

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins, for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})

	http.Handle("/", c.Handler(r))
	http.ListenAndServe(":8001", c.Handler(r))
}
