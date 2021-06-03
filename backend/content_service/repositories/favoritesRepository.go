package repositories

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"gorm.io/gorm"
)

type FavoritesRepository interface {
	GetAllCollections(context.Context, string) 			 		 ([]persistence.Collection, error)
	GetCollection(context.Context, string) 		 		 		 (persistence.Collection, error)

	GetAllFavorites(context.Context, string) 				 	 ([]persistence.Favorites, error)
	GetFavoritesFromCollection(context.Context, string)  		 ([]persistence.Favorites, error)

	CreateFavorite(context.Context, domain.FavoritesRequest) error
	RemoveFavorite(context.Context, domain.FavoritesRequest) error

	CreateCollection(context.Context, domain.Collection) error
	RemoveCollection(context.Context, string) error
}

type favoritesRepository struct {
	DB *gorm.DB
	contentRepository ContentRepository
}

func NewFavoritesRepo(db *gorm.DB) (*favoritesRepository, error) {
	if db == nil {
		panic("FavoritesRepository not created, gorm.DB is nil")
	}

	contentRepository, _ := NewContentRepo(db)

	return &favoritesRepository{ DB: db, contentRepository: contentRepository }, nil
}

func (repository *favoritesRepository) GetAllCollections(ctx context.Context, userId string) ([]persistence.Collection, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllCollections")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	collections := []persistence.Collection{}
	result := repository.DB.Where("user_id = ?", userId).Find(&collections)

	if result.Error != nil {
		return collections, result.Error
	}

	return collections, nil
}

func (repository *favoritesRepository) GetCollection(ctx context.Context, collectionId string) (persistence.Collection, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCollection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	collection := persistence.Collection{}
	result := repository.DB.Where("id = ?", collectionId).Find(&collection)

	if result.Error != nil {
		return collection, result.Error
	}

	return collection, nil
}

func (repository *favoritesRepository) GetAllFavorites(ctx context.Context, userId string) ([]persistence.Favorites, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllFavorites")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	favorites := []persistence.Favorites{}
	result := repository.DB.Where("user_id = ?", userId).Find(&favorites)

	if result.Error != nil {
		return favorites, result.Error
	}

	return favorites, nil
}

func (repository *favoritesRepository) GetFavoritesFromCollection(ctx context.Context, collectionId string) ([]persistence.Favorites, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetFavoritesFromCollection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	favorites := []persistence.Favorites{}
	result := repository.DB.Where("collection_id = ?", collectionId).Find(&favorites)

	if result.Error != nil {
		return favorites, result.Error
	}

	return favorites, nil
}

func (repository *favoritesRepository) CreateFavorite(ctx context.Context, favorites domain.FavoritesRequest) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFavorite")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var favoritesPers *persistence.Favorites
	favoritesPers = favoritesPers.ConvertToPersistence(favorites)

	// TODO Check if user can save that post
	result := repository.DB.Create(favoritesPers)

	if result.Error != nil || result.RowsAffected != 1 {
		return result.Error
	}

	return nil
}

func (repository *favoritesRepository) RemoveFavorite(ctx context.Context, favorites domain.FavoritesRequest) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveFavorite")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var favoritesPers *persistence.Favorites
	favoritesPers = favoritesPers.ConvertToPersistence(favorites)

	result := repository.DB.Delete(&favoritesPers)

	if result.Error != nil || result.RowsAffected != 1 {
		return result.Error
	}

	return nil
}

func (repository *favoritesRepository) CreateCollection(ctx context.Context, collection domain.Collection) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCollection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func (tx *gorm.DB) error{
		var collectionPers *persistence.Collection
		collectionPers = collectionPers.ConvertToPersistence(collection)

		// TODO Check if user can save that post
		result := repository.DB.Create(collectionPers)

		if result.Error != nil || result.RowsAffected != 1 {
			return result.Error
		}

		// Case: New collection was created upon saving post to favorites
		if len(collection.Posts) > 0 {
			for _, post := range collection.Posts {
				err := repository.contentRepository.CreatePost(ctx, &post)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil { return err }
	return nil
}

func (repository *favoritesRepository) RemoveCollection(ctx context.Context, collectionId string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveCollection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func (tx *gorm.DB) error{
		collectionPers := persistence.Collection{Id: collectionId}
		result := repository.DB.Delete(&collectionPers)

		if result.Error != nil || result.RowsAffected != 1 {
			return result.Error
		}

		collectionPosts, err := repository.contentRepository.GetCollectionsPosts(ctx, collectionId)
		if err != nil { return err }

		for _, post := range collectionPosts {
			err := repository.contentRepository.RemovePost(ctx, post.Id)
			if err != nil { return err }
		}

		return nil
	})

	if err != nil { return err }
	return nil
}
