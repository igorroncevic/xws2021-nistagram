package controllers

import (
	"encoding/json"
	"net/http"
	models2 "xws2021-nistagram/models"
	services2 "xws2021-nistagram/services"
	"xws2021-nistagram/util/errors"
)

type UserController struct {
	Service services2.UserService
}

func (controller *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := controller.Service.GetAllUsers()

	if err != nil {
		errors.WriteErrToClient(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models2.User

	json.NewDecoder(r.Body).Decode(&newUser)
	err := controller.Service.CreateUser(&newUser)
	if err != nil {
		errors.WriteErrToClient(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}
