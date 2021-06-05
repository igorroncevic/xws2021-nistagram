package setup

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

func CreateConnection(path string, dbName string, pw string) (neo4j.Driver, error) {
	driver, err := neo4j.NewDriver(path, neo4j.BasicAuth(dbName, pw, ""))
	if err != nil {
		panic("Database connection is not created!")
		return nil, err
	}
	log.Println("Application is successfully connected to Neo4j-GraphDb!")

	return driver, nil
}

func CloseConnection(driver neo4j.Driver) error {
	return driver.Close()
}

