package main

import (
	"github.com/gorilla/mux"
	"github.com/lytics/confl"
	"net/http"
	"xws2021-nistagram/util"
)

func main() {
	var dbConf util.DbConfig
	if _, err := confl.DecodeFile("dbconfig.conf", &dbConf); err != nil {
		panic(err)
	}

	dbPool := util.GetConnectionPool(dbConf)
	userController := util.GetUsersController(dbPool)

	r := mux.NewRouter()
	usersRouter := r.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("/hello", userController.GetAllUsers).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8001", r)
}
