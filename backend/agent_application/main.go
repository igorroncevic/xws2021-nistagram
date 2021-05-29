package main

import (
	"github.com/david-drvar/xws2021-nistagram/agent_application/util/setup"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	db := common.InitDatabase()
	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from agent application!"))
	}).Methods("GET")

	c := common.SetupCors()

	http.Handle("/", c.Handler(r))
	http.ListenAndServe(":8004", c.Handler(r))
}
