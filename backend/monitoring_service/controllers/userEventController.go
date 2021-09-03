package controllers

import (
	"github.com/igorroncevic/xws2021-nistagram/common/logger"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/services"
	"gorm.io/gorm"
)

type UserEventController struct{
	Service *services.UserEventService
	logger  *logger.Logger
}

func NewUserEventController(db *gorm.DB, logger *logger.Logger) UserEventController {
	userEventService, _ := services.NewUserEventService(db)

	return UserEventController{
		Service: userEventService,
		logger:  logger,
	}
}