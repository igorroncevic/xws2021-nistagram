package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type FavoritesService struct {
	favoritesRepository repositories.FavoritesRepository
	contentService      *PostService
}

func NewFavoritesService(db *gorm.DB) (*FavoritesService, error){
	favoritesRepository, err := repositories.NewFavoritesRepo(db)
	if err != nil {
		return nil, err
	}

	contentService, err := NewPostService(db)
	if err != nil {
		return nil, err
	}

	return &FavoritesService{
		favoritesRepository,
		contentService,
	}, err
}

func (service *FavoritesService) GetAllCollections(ctx context.Context, userId string) ([]domain.Collection, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllCollections")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCollections, err := service.favoritesRepository.GetAllCollections(ctx, userId)
	if err != nil {
		return []domain.Collection{}, err
	}

	collections := []domain.Collection{}
	for _, dbCollection := range dbCollections {
		dbPosts, err := service.favoritesRepository.GetFavoritesFromCollection(ctx, dbCollection.Id)
		if err != nil {
			return []domain.Collection{}, err
		}

		posts := []domain.ReducedPost{}
		for _, post := range dbPosts{
			// converted, err := service.contentService.GetReducedPostData(ctx, post.PostId)
			converted := domain.ReducedPost{
				Objava: domain.Objava{ Id: post.PostId },
			}
			if err != nil { return []domain.Collection{}, err }

			posts = append(posts, converted)
		}

		collection := dbCollection.ConvertToDomain(posts)
		collections = append(collections, collection)
	}

	dbUnclassified, err := service.favoritesRepository.GetUnclassifiedFavorites(ctx, userId)
	if err != nil { return []domain.Collection{}, err }

	unclassifiedPosts := []domain.ReducedPost{}
	for _, unclassifiedPost := range dbUnclassified{
		converted := domain.ReducedPost{
			Objava: domain.Objava{ Id: unclassifiedPost.Id },
		}
		if err != nil { return []domain.Collection{}, err }

		unclassifiedPosts = append(unclassifiedPosts, converted)
	}

	collections = append(collections, domain.Collection{
		Id:     "1",
		Name:   "No Collection",
		UserId: userId,
		Posts:  unclassifiedPosts,
	})

	return collections, nil
}
func (service *FavoritesService) GetCollection(ctx context.Context, collectionId string) (domain.Collection, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCollection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCollection, err := service.favoritesRepository.GetCollection(ctx, collectionId)
	if err != nil {
		return domain.Collection{}, err
	}

	dbPosts, err := service.favoritesRepository.GetFavoritesFromCollection(ctx, dbCollection.Id)
	if err != nil {
		return domain.Collection{}, err
	}

	posts := []domain.ReducedPost{}
	for _, post := range dbPosts{
		converted, err := service.contentService.GetReducedPostData(ctx, post.PostId)
		if err != nil { return domain.Collection{}, err }

		posts = append(posts, converted)
	}

	collection := dbCollection.ConvertToDomain(posts)

	return collection, nil
}

func (service *FavoritesService) GetUserFavorites(ctx context.Context, userId string) (domain.Favorites, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserFavorites")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	collections, err := service.GetAllCollections(ctx, userId)
	if err != nil {
		return domain.Favorites{}, nil
	}

	dbUnclassified, err := service.favoritesRepository.GetUnclassifiedFavorites(ctx, userId)
	if err != nil {
		return domain.Favorites{}, err
	}

	unclassified := []domain.ReducedPost{}
	for _, post := range dbUnclassified {
		converted, err := service.contentService.GetReducedPostData(ctx, post.Id)
		if err != nil {
			return domain.Favorites{}, err
		}
		unclassified = append(unclassified, converted)
	}

	favorites := domain.Favorites{
		UserId:       userId,
		Collections:  collections,
		Unclassified: unclassified,
	}

	return favorites, nil
}

func (service *FavoritesService) CreateFavorite(ctx context.Context, favoritesRequest domain.FavoritesRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFavorite")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.favoritesRepository.CreateFavorite(ctx, favoritesRequest)
	if err != nil{
		return err
	}

	return  nil
}
func (service *FavoritesService) RemoveFavorite(ctx context.Context, favoritesRequest domain.FavoritesRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveFavorite")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.favoritesRepository.RemoveFavorite(ctx, favoritesRequest)
	if err != nil {
		return err
	}

	return nil
}

func (service *FavoritesService) CreateCollection(ctx context.Context, collection domain.Collection) (domain.Collection, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCollection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCollection, err := service.favoritesRepository.CreateCollection(ctx, collection)
	if err != nil {
		return domain.Collection{}, err
	}

	collection.Id = dbCollection.Id

	return collection, nil
}
func (service *FavoritesService) RemoveCollection(ctx context.Context, collectionId string, userId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveCollection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.favoritesRepository.RemoveCollection(ctx, collectionId, userId)
	if err != nil {
		return err
	}

	return nil
}