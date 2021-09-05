package controllers

import (
	"context"
	"encoding/json"
	"github.com/igorroncevic/xws2021-nistagram/common"
	"github.com/igorroncevic/xws2021-nistagram/common/logger"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/services"
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/util"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type UserEventController struct{
	Service    *services.UserEventService
	logger     *logger.Logger
	jwtManager *common.JWTManager
}

func NewUserEventController(db *gorm.DB, logger *logger.Logger, jwtManager *common.JWTManager) UserEventController {
	userEventService, _ := services.NewUserEventService(db)

	return UserEventController{
		Service: userEventService,
		logger:  logger,
		jwtManager: jwtManager,
	}
}

func (c UserEventController) GetUsersActivity(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAds")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	// TODO Implement extracting from JWT
	header := r.Header.Get("Authorization")
	token := strings.Split(header, " ")[1]

	claims, err := c.jwtManager.ExtractClaims(token)
	if err != nil || claims.UserId == "" {
		c.logger.ToStdoutAndFile("GetUsersActivity", "Failed to authenticate", logger.Error)
		util.WriteErrToClient(w, err)
		return
	}

	events, err := c.Service.GetUsersActivity(ctx, claims.UserId, claims.Email)
	if err != nil {
		c.logger.ToStdoutAndFile("GetUsersActivity", "Failed to gather all activity", logger.Error)
		util.WriteErrToClient(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}