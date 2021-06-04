package main

import (
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/util/setup"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main(){
	driver, _ := setup.CreateConnection("bolt://localhost:7687", "neo4j", "root")
	defer setup.CloseConnection(driver)

	err := CreateUniqueConstraint(driver)
	if err != nil {
		panic("Could not create unique constraint!")
	}

	setup.GRPCServer(driver)

}


func CreateUniqueConstraint(driver neo4j.Driver) error {

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run(
			"CREATE CONSTRAINT constraint_id IF NOT EXISTS ON (u:User) ASSERT u.id IS UNIQUE",
			map[string]interface{}{

			})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return  err
	}
	return  nil
}