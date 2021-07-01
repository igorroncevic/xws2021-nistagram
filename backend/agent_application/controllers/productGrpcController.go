package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/agent_application/services"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductGrpcController struct {
	service    *services.ProductService
	jwtManager *common.JWTManager
	logger     *logger.Logger
}

func NewProductController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger) (*ProductGrpcController, error) {
	service, err := services.NewProductService(db)
	if err != nil {
		return nil, err
	}

	return &ProductGrpcController{
		service,
		jwtManager,
		logger,
	}, nil
}

func (s *ProductGrpcController) CreateProduct(ctx context.Context, in *protopb.Product) (*protopb.EmptyResponseAgent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var product *persistence.Product
	product = product.ConvertFromGrpc(in)
	if product == nil {
		return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, "cannot create product")
	}

	err := s.service.CreateProduct(ctx, *product)
	if err != nil {
		return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &protopb.EmptyResponseAgent{}, nil
}
