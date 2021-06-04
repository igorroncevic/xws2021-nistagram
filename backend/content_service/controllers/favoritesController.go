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

type FavoritesGrpcController struct {
	service *services.FavoritesService
}

func NewFavoritesController(db *gorm.DB) (*FavoritesGrpcController, error) {
	service, err := services.NewFavoritesService(db)
	if err != nil {
		return nil, err
	}

	return &FavoritesGrpcController{
		service,
	}, nil
}

func (c *FavoritesGrpcController) GetAllCollections (ctx context.Context, in *contentpb.RequestId) (*contentpb.CollectionsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllCollections")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	collections, err := c.service.GetAllCollections(ctx, in.Id)
	if err != nil {
		return &contentpb.CollectionsArray{}, status.Errorf(codes.Unknown, "could not retrieve collections")
	}

	grpcCollections := []*contentpb.Collection{}
	for _, collection := range collections{
		grpcCollections = append(grpcCollections, collection.ConvertToGrpc())
	}

	return &contentpb.CollectionsArray{
		Collections: grpcCollections,
	}, nil
}

func (c *FavoritesGrpcController) GetCollection (ctx context.Context, in *contentpb.RequestId) (*contentpb.Collection, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCollection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	collection, err := c.service.GetCollection(ctx, in.Id)
	if err != nil || collection.Id == "" {
		return &contentpb.Collection{}, status.Errorf(codes.Unknown, "could not retrieve collection")
	}

	grpcCollection := collection.ConvertToGrpc()

	return grpcCollection, nil
}

func (c *FavoritesGrpcController) GetUserFavorites (ctx context.Context, in *contentpb.RequestId) (*contentpb.Favorites, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserFavorites")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	favorites, err := c.service.GetUserFavorites(ctx, in.Id)
	if err != nil {
		return &contentpb.Favorites{}, status.Errorf(codes.Unknown, "could not retrieve favorites")
	}

	grpcFavorites := favorites.ConvertToGrpc()

	return grpcFavorites, nil
}

func (c *FavoritesGrpcController) CreateFavorite (ctx context.Context, in *contentpb.FavoritesRequest) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFavorite")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var favoritesRequest domain.FavoritesRequest
	favoritesRequest = favoritesRequest.ConvertFromGrpc(in)

	err := c.service.CreateFavorite(ctx, favoritesRequest)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not create favorite")
	}

	return &contentpb.EmptyResponse{}, nil
}

func (c *FavoritesGrpcController) RemoveFavorite (ctx context.Context, in *contentpb.FavoritesRequest) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveFavorite")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var favoritesRequest domain.FavoritesRequest
	favoritesRequest = favoritesRequest.ConvertFromGrpc(in)

	err := c.service.RemoveFavorite(ctx, favoritesRequest)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not remove favorite")
	}

	return &contentpb.EmptyResponse{}, nil
}

func (c *FavoritesGrpcController) CreateCollection (ctx context.Context, in *contentpb.Collection) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCollection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var collection domain.Collection
	collection = collection.ConvertFromGrpc(in)

	err := c.service.CreateCollection(ctx, collection)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not create collection")
	}

	return &contentpb.EmptyResponse{}, nil
}

func (c *FavoritesGrpcController) RemoveCollection (ctx context.Context, in *contentpb.RequestId) (*contentpb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveFavorite")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := c.service.RemoveCollection(ctx, in.Id)
	if err != nil {
		return &contentpb.EmptyResponse{}, status.Errorf(codes.Unknown, "could not remove collection")
	}

	return &contentpb.EmptyResponse{}, nil
}