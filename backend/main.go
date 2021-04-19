package main

import (
	"github.com/gorilla/mux"
	"github.com/lytics/confl"
	"net/http"
	"xws2021-nistagram/util"
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
	usersRouter.HandleFunc("", userController.CreateUser).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8001", r)
}
