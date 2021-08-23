package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/igorroncevic/xws2021-nistagram/common"
	"github.com/igorroncevic/xws2021-nistagram/common/kafka_util"
	"github.com/igorroncevic/xws2021-nistagram/common/logger"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/controllers"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/model"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/util"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/util/setup"
	"net/http"
	"os"
	"time"
)

func main(){
	if os.Getenv("Docker_env") == "" {
		SetupEnvVariables()
	}
	db := common.InitDatabase(common.MonitoringDatabase)
	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}

	performanceConsumer := kafka_util.NewConsumer(kafka_util.RegularConsumerMaxWait, kafka_util.ExampleGroupId, kafka_util.PerformanceTopic)
	userEventsConsumer := kafka_util.NewConsumer(kafka_util.RegularConsumerMaxWait, kafka_util.ExampleGroupId, kafka_util.UserEventsTopic)
	customLogger := logger.NewLogger()

	performanceController := controllers.NewPerformanceController(db, customLogger)
	userEventController := controllers.NewUserEventController(db, customLogger)

	r := mux.NewRouter()

	performanceRouter := r.PathPrefix("/performance").Subrouter()
	performanceRouter.HandleFunc("", performanceController.GetAllStats).Methods("GET")

	go performanceConsumer.Consume(func(id string, timestamp time.Time, message map[string]interface{}) error {
		converted := kafka_util.ConvertToPerformanceMessage(message)

		// Ignoring 1xx, 2xx and 3xx status codes, we are only interested in errors
		if util.GetFirstDigit(converted.Status) != 4 && util.GetFirstDigit(converted.Status) != 5 {
			return nil
		}

		persistMessage := model.ConvertPerformanceMessageToPersistence(id, timestamp, converted)
		return performanceController.Service.SaveEntry(context.Background(), persistMessage)
	})
	defer performanceConsumer.Close()

	go userEventsConsumer.Consume(func(id string, timestamp time.Time, message map[string]interface{}) error {
		converted := kafka_util.ConvertToUserEventMessage(message)

		persistMessage := model.ConvertUserEventMessageToPersistence(id, timestamp, converted)
		return userEventController.Service.SaveEntry(context.Background(), persistMessage)
	})
	defer userEventsConsumer.Close()

	customLogger.ToStdoutAndFile("Monitoring Service", "Starting service...", logger.Info)

	http.Handle("/", r)
	http.ListenAndServe(":8006", r)
}

func SetupEnvVariables() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", common.MonitoringDatabaseName)
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PW", "root")
}