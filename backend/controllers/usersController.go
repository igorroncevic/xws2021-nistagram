package controllers

import (
	"encoding/json"
	"net/http"
	"xws2021-nistagram/backend/models"
	"xws2021-nistagram/backend/models/dtos"
	"xws2021-nistagram/backend/services"
	"xws2021-nistagram/backend/util/errors"
)

type UserController struct {
	Service services.UserService
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
	var newUser models.User

	json.NewDecoder(r.Body).Decode(&newUser)
	err := controller.Service.CreateUser(&newUser)
	if err != nil {
		errors.WriteErrToClient(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (controller *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData dtos.LoginDTO

	json.NewDecoder(r.Body).Decode(&loginData)
	err := controller.Service.LoginUser(loginData)
	if err != nil {
		errors.WriteErrToClient(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
