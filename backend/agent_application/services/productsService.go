package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/agent_application/repositories"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"gorm.io/gorm"
)

type ProductService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(db *gorm.DB) (*ProductService, error) {
	productRepository, err := repositories.NewProductRepo(db)

	return &ProductService{
		productRepository,
	}, err
}

func (service *ProductService) CreateProduct(ctx context.Context, product persistence.Product) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserInAgentApp")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return err
	}

	return nil
}
