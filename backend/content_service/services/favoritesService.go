package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type FavoritesService struct {
	favoritesRepository repositories.FavoritesRepository
	contentRepository repositories.ContentRepository
}

func NewFavoritesService(db *gorm.DB) (*FavoritesService, error){
	favoritesRepository, err := repositories.NewFavoritesRepo(db)
	if err != nil {
		return nil, err
	}

	contentRepository, err := repositories.NewContentRepo(db)
	if err != nil {
		return nil, err
	}

	return &FavoritesService{
		favoritesRepository,
		contentRepository,
	}, err
}

func (service *FavoritesService) GetAllCollections(context.Context, string) ([]domain.Collection, error) {

	return []domain.Collection{}, nil
}
func (service *FavoritesService) GetCollection(ctx context.Context, collectionId string) (domain.Collection, error) {

	return domain.Collection{}, nil
}

func (service *FavoritesService) GetAllFavorites(ctx context.Context, id string) ([]domain.Favorites, error) {

	return []domain.Favorites{}, nil
}
func (service *FavoritesService) GetFavoritesFromCollection(ctx context.Context, collectionId string) ([]domain.Favorites, error) {

	return []domain.Favorites{}, nil
}

func (service *FavoritesService) CreateFavorite(ctx context.Context, favoritesRequest domain.FavoritesRequest) error {

	return  nil
}
func (service *FavoritesService) RemoveFavorite(ctx context.Context, favoritesRequest domain.FavoritesRequest) error {

	return nil
}

func (service *FavoritesService) CreateCollection(ctx context.Context, collection domain.Collection) error {

	return nil
}
func (service *FavoritesService) RemoveCollection(ctx context.Context, collectionId string) error {

	return nil
}