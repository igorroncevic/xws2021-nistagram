package main

import (
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"net/http"
)

func main(){
/*	db := common.InitDatabase()
	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}*/
	result, err := helloWorld("bolt://localhost:7687", "neo4j", "root")
	if err != nil {
		panic(err)
	}
	log.Println(result)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from recommendation service!"))
	}).Methods("GET")

	c := common.SetupCors()

	http.Handle("/", c.Handler(r))
	http.ListenAndServe(":8005", c.Handler(r))
}

func helloWorld(uri, username, password string) (string, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return "", err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (n:Greeting) RETURN n.message + ' ' + id(n)", map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		if result.Next()  {
			return result.Record().Values[0], nil
		}
		result, err = transaction.Run(
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]interface{}{"message": "hello, world"})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}

	return greeting.(string), nil
}