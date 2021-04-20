package main

import (
	"github.com/gorilla/mux"
	"github.com/lytics/confl"
	"github.com/rs/cors"
	"net/http"
	"strings"
	"xws2021-nistagram/backend/util"
	"xws2021-nistagram/backend/util/auth"
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

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", userController.CreateUser).Methods("POST")
	authRouter.HandleFunc("/login", userController.LoginUser).Methods("POST")

	usersRouter.Use(authMiddleware) // Authenticate user's JWT before letting them access these endpoints

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins, for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})

	http.Handle("/", c.Handler(r))
	http.ListenAndServe(":8001", c.Handler(r))
}

func authMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.String(), "favicon.ico") {
			// Allow favicon.ico to load
			next.ServeHTTP(w, r)
		}

		authHeader := r.Header.Get("Authorization")
		splitHeader := strings.Split(authHeader, " ")
		if len(splitHeader) != 2{
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jwtString := splitHeader[1]
		status, err := auth.ValidateJWT(jwtString)

		if err != nil{
			w.WriteHeader(status)
			return
		}else{
			next.ServeHTTP(w, r)
		}
	})
}
