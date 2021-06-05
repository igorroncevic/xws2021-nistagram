package controllers

import (
	"context"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	contentpb "github.com/david-drvar/xws2021-nistagram/content_service/proto"
	"github.com/david-drvar/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type HighlightGrpcController struct {
	service *services.HighlightService
}

func NewHighlightController(db *gorm.DB) (*HighlightGrpcController, error) {
	service, err := services.NewHighlightService(db)
	if err != nil {
		return nil, err
	}

	return &HighlightGrpcController{
		service,
	}, nil
}

func (c *HighlightGrpcController) GetAllHighlights (ctx context.Context, in *contentpb.RequestId) (*contentpb.HighlightsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHighlights")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	highlights, err := c.service.GetAllHighlights(ctx, in.Id)
	if err != nil {
		return &contentpb.HighlightsArray{}, status.Errorf(codes.Unknown, "could not retrieve highlights")
	}

	grpcHighlights := []*contentpb.Highlight{}
	for _, highlight := range highlights{
		grpcHighlights = append(grpcHighlights, highlight.ConvertToGrpc())
	}

	return &contentpb.HighlightsArray{
		Highlights: grpcHighlights,
	}, nil
}

func (c *HighlightGrpcController) GetHighlight (ctx context.Context, in *contentpb.RequestId) (*contentpb.Highlight, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetHighlight")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	highlight, err := c.service.GetHighlight(ctx, in.Id)
	if err != nil || highlight.Id == "" {
		return &contentpb.Highlight{}, status.Errorf(codes.Unknown, "could not retrieve highlight")
	}

	grpcHighlight := highlight.ConvertToGrpc()

	return grpcHighlight, nil
}

func (c *HighlightGrpcController) CreateHighlightStory(ctx context.Context, in *contentpb.HighlightRequest) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHighlightStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var highlightRequest *domain.HighlightRequest
	highlightRequest = highlightRequest.ConvertFromGrpc(in)

	err := c.service.CreateHighlightStory(ctx, *highlightRequest)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not create story from highlight")
	}

	return &contentpb.EmptyResponse{}, nil
}

func (c *HighlightGrpcController) RemoveHighlightStory(ctx context.Context, in *contentpb.HighlightRequest) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveHighlightStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var highlightRequest *domain.HighlightRequest
	highlightRequest = highlightRequest.ConvertFromGrpc(in)

	err := c.service.RemoveHighlightStory(ctx, *highlightRequest)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not remove story from highlight")
	}

	return &contentpb.EmptyResponse{}, nil
}

func (c *HighlightGrpcController) CreateHighlight (ctx context.Context, in *contentpb.Highlight) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHighlight")
	defer span.Finish()
	contextMetadata, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(contextMetadata)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var collection *domain.Highlight
	collection = collection.ConvertFromGrpc(in)

	err := c.service.CreateHighlight(ctx, *collection)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not create highlight")
	}

	return &contentpb.EmptyResponse{}, nil
}

func (c *HighlightGrpcController) RemoveHighlight (ctx context.Context, in *contentpb.RequestId) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveHighlight")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := c.service.RemoveHighlight(ctx, in.Id)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not remove highlight")
	}

	return &contentpb.EmptyResponse{}, nil
}