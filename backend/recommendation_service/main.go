package main

import (
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/util/setup"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	db := common.InitDatabase(common.RecommendationDatabase)
	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from recommendation service!"))
	}).Methods("GET")

	c := common.SetupCors()

	http.Handle("/", c.Handler(r))
	http.ListenAndServe(":8005", c.Handler(r))
}
