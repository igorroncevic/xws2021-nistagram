package main

import (
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/setup"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db := common.InitDatabase()
	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}
	userController := setup.GetUsersController(db)
	privacyController := setup.GetPrivacyController(db)

	r := mux.NewRouter()

	usersRouter := r.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", userController.GetAllUsers).Methods("GET")
	usersRouter.HandleFunc("/update_profile", userController.UpdateUserProfile).Methods("POST")

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", userController.CreateUser).Methods("POST")
	authRouter.HandleFunc("/login", userController.LoginUser).Methods("POST")

	privacyRouter := r.PathPrefix("/privacy").Subrouter()
	privacyRouter.HandleFunc("", privacyController.CreatePrivacy).Methods("Post")

	usersRouter.Use(common.AuthMiddleware) // Authenticate user's JWT before letting them access these endpoints

	c := common.SetupCors()

	http.Handle("/", c.Handler(r))
	http.ListenAndServe(":8001", c.Handler(r))
}
