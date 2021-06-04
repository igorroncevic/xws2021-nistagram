package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	contentpb "github.com/david-drvar/xws2021-nistagram/content_service/proto"
	"github.com/david-drvar/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type StoryGrpcController struct {
	service *services.StoryService
}

func NewStoryController(db *gorm.DB) (*StoryGrpcController, error) {
	service, err := services.NewStoryService(db)
	if err != nil {
		return nil, err
	}

	return &StoryGrpcController{
		service:  service,
	}, nil
}

func (c *StoryGrpcController) GetAllHomeStories(ctx context.Context, in *contentpb.EmptyRequest) (*contentpb.StoriesArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHomeStories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	stories, err := c.service.GetAllHomeStories(ctx)

	if err != nil{
		return &contentpb.StoriesArray{
			Stories: []*contentpb.Story{},
		}, status.Errorf(codes.Unknown, err.Error())
	}

	responseStories := []*contentpb.Story{}
	for _, story := range stories {
		responseStories = append(responseStories, story.ConvertToGrpc())
	}

	return &contentpb.StoriesArray{
		Stories: responseStories,
	}, nil
}

func (c *StoryGrpcController) CreateStory(ctx context.Context, in *contentpb.Story) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var story *domain.Story
	story = story.ConvertFromGrpc(in)

	err := c.service.CreateStory(ctx, story)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not create story")
	}

	return &contentpb.EmptyResponse{}, nil
}

func (c *StoryGrpcController) GetStoryById(ctx context.Context, in *contentpb.RequestId) (*contentpb.Story, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetStoryById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	story, err := c.service.GetStoryById(ctx, in.Id)

	if err != nil { return &contentpb.Story{}, status.Errorf(codes.Unknown, err.Error()) }

	grpcStory := story.ConvertToGrpc()

	return grpcStory, nil
}

func (c *StoryGrpcController) RemoveStory(ctx context.Context, in *contentpb.RequestId) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := c.service.RemoveStory(ctx, in.Id)

	if err != nil{
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &contentpb.EmptyResponse{}, nil
}