package main

import (
	"github.com/gorilla/mux"
	"github.com/lytics/confl"
	"github.com/rs/cors"
	"net/http"
	"xws2021-nistagram/backend/util"
)

func main() {
	var dbConf util.DbConfig
	if _, err := confl.DecodeFile("./backend/dbconfig.conf", &dbConf); err != nil {
		panic(err)
	}

	dbPool := util.GetConnectionPool(dbConf)
	userController := util.GetUsersController(dbPool)

	r := mux.NewRouter()

	usersRouter := r.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("", userController.GetAllUsers).Methods("GET")
	usersRouter.HandleFunc("/register", userController.CreateUser).Methods("POST")
	usersRouter.HandleFunc("/login", userController.LoginUser).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins, for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})

	http.Handle("/", c.Handler(r))
	http.ListenAndServe(":8001", c.Handler(r))
}
