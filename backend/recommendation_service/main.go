package main

import (
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/util/setup"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"os"
	"time"
)

func main(){
	if os.Getenv("Docker_env") == "" {
		SetupEnvVariables()
	}

	time.Sleep(30*time.Second)

	driver, _ := setup.CreateConnection(os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_PW"))
	defer setup.CloseConnection(driver)
	for  i := 1; i < 5; i++ {
		duration := 2 << i - 1
		time.Sleep(time.Duration(duration)*time.Second)
		err := CreateUniqueConstraint(driver)
		if err != nil {
			log.Println(err)
			log.Println("AJDEEEE")
			log.Println("Retrying . . . ")
			continue
		}else {
			log.Println("Success!")
		}
		break
	}



	setup.GRPCServer(driver)

}

func SetupEnvVariables() {
	os.Setenv("DB_HOST", "bolt://localhost:7687")
	os.Setenv("DB_NAME", common.RecommendationDatabaseName)
	os.Setenv("DB_PW", "root")
}



func CreateUniqueConstraint(driver neo4j.Driver) error {
	log.Println("Nece error da ispise")

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run(
			"CREATE CONSTRAINT constraint_id IF NOT EXISTS ON (u:User) ASSERT u.id IS UNIQUE",
			map[string]interface{}{

			})
		if err != nil {
			log.Println(err)
			log.Println("Nece error da ispise")
			return nil, err
		}
		log.Println("A ja ovde baka....")
		return nil, nil
	})
	if err != nil {
		return  err
	}
	return  nil
}