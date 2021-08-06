package setup

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"os"
)

func CreateConnection(path string, dbName string, pw string) (neo4j.Driver, error) {
	driver, err := neo4j.NewDriver(path, neo4j.BasicAuth(dbName, pw, ""))
	log.Println("Environment variables: " + os.Getenv("DB_HOST") + os.Getenv("DB_NAME") + os.Getenv("DB_PW"))

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
