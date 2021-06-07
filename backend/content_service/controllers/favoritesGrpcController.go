package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type FavoritesGrpcController struct {
	service     *services.FavoritesService
	postService *services.PostService
	jwtManager *common.JWTManager
}

func NewFavoritesController(db *gorm.DB, jwtManager *common.JWTManager) (*FavoritesGrpcController, error) {
	service, err := services.NewFavoritesService(db)
	if err != nil {
		return nil, err
	}

	postService, err := services.NewPostService(db)
	if err != nil {
		return nil, err
	}

	return &FavoritesGrpcController{
		service,
		postService,
		jwtManager,
	}, nil
}

func (c *FavoritesGrpcController) GetAllCollections(ctx context.Context, in *protopb.RequestId) (*protopb.CollectionsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllCollections")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.CollectionsArray{}, status.Errorf(codes.Unknown, err.Error())
	}  else if claims.UserId == "" {
		return &protopb.CollectionsArray{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}  else if claims.UserId != in.Id {
		return &protopb.CollectionsArray{}, status.Errorf(codes.Unknown, "cannot get another user's collections")
	}

	collections, err := c.service.GetAllCollections(ctx, in.Id)
	if err != nil {
		return &protopb.CollectionsArray{}, status.Errorf(codes.Unknown, "could not retrieve collections")
	}

	grpcCollections := []*protopb.Collection{}
	for _, collection := range collections {
		grpcCollections = append(grpcCollections, collection.ConvertToGrpc())
	}

	return &protopb.CollectionsArray{
		Collections: grpcCollections,
	}, nil
}

func (c *FavoritesGrpcController) GetCollection(ctx context.Context, in *protopb.RequestId) (*protopb.Collection, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCollection")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.Collection{}, status.Errorf(codes.Unknown, "could not retrieve collection")
	}else if claims.UserId == "" {
		return &protopb.Collection{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}  else if claims.UserId != in.Id {
		return &protopb.Collection{}, status.Errorf(codes.Unknown, "cannot get another user's collection")
	}

	collection, err := c.service.GetCollection(ctx, in.Id)
	if err != nil || collection.Id == "" {
		return &protopb.Collection{}, status.Errorf(codes.Unknown, "could not retrieve collection")
	}

	grpcCollection := collection.ConvertToGrpc()

	return grpcCollection, nil
}

func (c *FavoritesGrpcController) GetUserFavorites(ctx context.Context, in *protopb.RequestId) (*protopb.Favorites, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserFavorites")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.Favorites{}, status.Errorf(codes.Unknown, "could not retrieve favorites")
	}else if claims.UserId == "" {
		return &protopb.Favorites{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}  else if claims.UserId != in.Id {
		return &protopb.Favorites{}, status.Errorf(codes.Unknown, "cannot get another user's favorites")
	}

	favorites, err := c.service.GetUserFavorites(ctx, in.Id)
	if err != nil {
		return &protopb.Favorites{}, status.Errorf(codes.Unknown, "could not retrieve favorites")
	}

	grpcFavorites := favorites.ConvertToGrpc()

	return grpcFavorites, nil
}

func (c *FavoritesGrpcController) CreateFavorite(ctx context.Context, in *protopb.FavoritesRequest) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFavorite")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create favorite")
	}else if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}  else if claims.UserId != in.UserId {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot create favorite for another user")
	}

	post, err := c.postService.GetReducedPostData(ctx, in.PostId)
	if err != nil { return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot create favorite") }

	if post.UserId != claims.UserId{
		following, err := grpc_common.CheckFollowInteraction(ctx, post.UserId, claims.UserId)
		if err != nil {
			return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot tag selected users")
		}

		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, post.UserId)
		if err != nil {
			return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
		}

		isBlocked, err := grpc_common.CheckIfBlocked(ctx, post.UserId, claims.UserId)
		if err != nil {
			return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
		}

		// If used is blocked or his profile is private and did not approve your request
		if isBlocked || (!isPublic && !following.IsApprovedRequest ) {
			return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot tag selected users")
		}
	}

	var favoritesRequest domain.FavoritesRequest
	favoritesRequest = favoritesRequest.ConvertFromGrpc(in)

	err = c.service.CreateFavorite(ctx, favoritesRequest)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create favorite")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *FavoritesGrpcController) RemoveFavorite(ctx context.Context, in *protopb.FavoritesRequest) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveFavorite")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not remove favorite")
	}else if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}  else if claims.UserId != in.UserId {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot remove favorite for another user")
	}

	var favoritesRequest domain.FavoritesRequest
	favoritesRequest = favoritesRequest.ConvertFromGrpc(in)

	err = c.service.RemoveFavorite(ctx, favoritesRequest)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not remove favorite")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *FavoritesGrpcController) CreateCollection(ctx context.Context, in *protopb.Collection) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCollection")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create collection")
	}else if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}  else if claims.UserId != in.UserId {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot create collection for another user")
	}

	var collection domain.Collection
	collection = collection.ConvertFromGrpc(in)

	err = c.service.CreateCollection(ctx, collection)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create collection")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *FavoritesGrpcController) RemoveCollection(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveFavorite")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create favorite")
	}else if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}

	err = c.service.RemoveCollection(ctx, in.Id, claims.UserId)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not remove collection")
	}

	return &protopb.EmptyResponseContent{}, nil
}
