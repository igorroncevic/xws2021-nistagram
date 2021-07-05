package repository

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/chat_service/model"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"gorm.io/gorm"
)

type MessageRepository interface {
	SaveMessage(context.Context, model.Message) error
	DeleteMessage( context.Context, string) error
	GetMessagesForChatRoom(context.Context, string) ([]model.Message, error)
	GetChatRoomsForUser(context.Context, string) ([]model.ChatRoom, error)
	CreateChatRoom(context.Context, model.ChatRoom) error
}

type messageRepository struct {
	db *gorm.DB

}

func NewMessageRepository (db *gorm.DB) (MessageRepository, error) {
	if db == nil {
		panic("MessageRepository not created, gorm.DB is nil")
	}
	return &messageRepository{db: db}, nil
}

func (repo *messageRepository) 	SaveMessage(ctx context.Context, message model.Message) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repo.db.Create(&message)
	if result.Error != nil {
		return errors.New("Could not save message!")
	}

	return nil
}

func (repo *messageRepository) DeleteMessage( ctx context.Context, id string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := repo.db.Where("id = ?", id).Delete(&model.Message{})
	if result.Error != nil {
		return errors.New("Could not delete message")
	}

	return nil
}

func (repo *messageRepository) 	CreateChatRoom(ctx context.Context, room model.ChatRoom) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var checkRoom *model.ChatRoom
	result := repo.db.Where("id = ?", room.Person1+room.Person2).Find(&checkRoom)
	if result.Error != nil {
		return errors.New("Could not load chatRoom")
	}else if checkRoom.Id != "" {
		return errors.New("Room already exists")
	}

	result = repo.db.Where("id = ?", room.Person2+room.Person1).Find(&checkRoom)
	if result.Error != nil {
		return errors.New("Could not load chatRoom")
	}else if checkRoom.Id != "" {
		return errors.New("Room already exists")
	}
	room.Id = room.Person1 + room.Person2
	result = repo.db.Create(&room)
	if result.Error != nil {
		return errors.New("Could not create chat room!")
	}else if result.RowsAffected == 0 {
		return errors.New("Could not create chat room!")
	}

	return nil

}

func (repo *messageRepository) GetChatRoomsForUser(ctx context.Context, userId string) ([]model.ChatRoom, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetChatRoomsForUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var rooms1 []model.ChatRoom
	result := repo.db.Where("person1 = ?", userId).Find(&rooms1)
	if result.Error != nil {
		return nil, errors.New("Could not load rooms for user")
	}
	var rooms2 []model.ChatRoom
	result = repo.db.Where("person2 = ?", userId).Find(&rooms2)
	if result.Error != nil {
		return nil, errors.New("Could not load rooms for user")
	}

	return append(rooms1, rooms2...), nil
}

func (repo *messageRepository) GetMessagesForChatRoom(ctx context.Context, roomId string) ([]model.Message, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetMessagesForChatRoom")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var messages []model.Message
	var room model.ChatRoom
	result := repo.db.Where("id = ?", roomId).Find(&room)
	if result.Error != nil {
		return []model.Message{}, errors.New("Could not load messages")
	}

	var messages1 []model.Message
	result = repo.db.Where("sender_id = ? ", room.Person1).Where("receiver_id = ?", room.Person2).Find(&messages1)
	if result.Error != nil {
		return nil, errors.New("Could not load messages")
	}
	var messages2 []model.Message
	result = repo.db.Where("sender_id = ? ", room.Person2).Where("receiver_id = ?", room.Person1).Find(&messages2)
	if result.Error != nil {
		return nil, errors.New("Could not load messages")
	}

	messages = append(messages, messages1...)
	messages = append(messages, messages2...)

	return messages, nil

}

