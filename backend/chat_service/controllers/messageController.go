package controllers

import (
	"context"
	"encoding/json"
	"github.com/david-drvar/xws2021-nistagram/chat_service/model"
	"github.com/david-drvar/xws2021-nistagram/chat_service/service"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type MessageController struct {
	Service *service.MessageService
}

func NewMessageController(db *gorm.DB) (*MessageController, error) {
	service, err := service.NewMessageService(db)
	if err != nil {
		return nil, err
	}

	return &MessageController{
		Service: service,
	}, err
}

func (c MessageController) SaveMessage(w http.ResponseWriter, r *http.Request)  {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var message model.Message
	json.NewDecoder(r.Body).Decode(&message)

	err := c.Service.SaveMessage(ctx, message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	json.NewEncoder(w).Encode(message)

}

func (c MessageController)	DeleteMessage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	id := mux.Vars(r)["id"]

	err := c.Service.DeleteMessage(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}


}

func (c MessageController) GetMessagesForChatRoom(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "GetMessagesForChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	roomId := mux.Vars(r)["roomId"]

	messages, err := c.Service.GetMessagesForChatRoom(ctx, roomId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	json.NewEncoder(w).Encode(messages)
}

func (c MessageController) GetChatRoomsForUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "GetChatRoomsForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	userId := mux.Vars(r)["userId"]

	rooms, err := c.Service.GetChatRoomsForUser(ctx, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	json.NewEncoder(w).Encode(rooms)

}

func (c MessageController) CreateChatRoom(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var room model.ChatRoom
	json.NewDecoder(r.Body).Decode(&room)

	err := c.Service.CreateChatRoom(ctx, room)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	json.NewEncoder(w).Encode(room)

}




