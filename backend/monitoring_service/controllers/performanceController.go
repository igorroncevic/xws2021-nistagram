package controllers

import (
	"context"
	"encoding/json"
	"github.com/igorroncevic/xws2021-nistagram/common/logger"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/services"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/util"
	"gorm.io/gorm"
	"net/http"
)

type PerformanceController struct{
	Service *services.PerformanceService
	logger  *logger.Logger
}

func NewPerformanceController(db *gorm.DB, logger *logger.Logger) PerformanceController {
	performanceService, _ := services.NewPerformanceService(db)

	return PerformanceController{
		Service: performanceService,
		logger:  logger,
	}
}

func (controller *PerformanceController) GetAllStats(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContext(ctx, "GetAllStats")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	controller.logger.ToStdoutAndFile("GetAllStats", "Requesting all stats", logger.Info)

	stats, err := controller.Service.GetAllStats(ctx)
	if err != nil {
		controller.logger.ToStdoutAndFile("GetAllStats", "Failed to gather all stats", logger.Error)
		util.WriteErrToClient(w, err)
		return
	}

	controller.logger.ToStdoutAndFile("GetAllStats", "Successfully gathered all stats", logger.Info)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}