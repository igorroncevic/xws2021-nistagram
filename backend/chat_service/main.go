package main

import (
	"github.com/david-drvar/xws2021-nistagram/chat_service/util/setup"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	db := common.InitDatabase("")
	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from chat service!"))
	}).Methods("GET")

	c := common.SetupCors()

	http.Handle("/", c.Handler(r))
	http.ListenAndServe(":8003", c.Handler(r))
}
