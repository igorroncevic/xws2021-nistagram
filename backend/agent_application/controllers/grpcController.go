package controllers

import (
	"context"
	"github.com/igorroncevic/xws2021-nistagram/common"
	"github.com/igorroncevic/xws2021-nistagram/common/logger"
	protopb "github.com/igorroncevic/xws2021-nistagram/common/proto"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	otgo "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"io"
)

type Server struct {
	protopb.UnimplementedAgentServer
	userController     *UserGrpcController
	productController  *ProductGrpcController
	campaignController *CampaignGrpcController
	tracer             otgo.Tracer
	closer             io.Closer
}

func NewServer(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger) (*Server, error) {
	newUserController, _ := NewUserController(db, jwtManager, logger)
	newProductController, _ := NewProductController(db, jwtManager, logger)
	newCampaignController, _ := NewCampaignController(db, jwtManager, logger)

	tracer, closer := tracer.Init("agentService")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		userController:     newUserController,
		productController:  newProductController,
		campaignController: newCampaignController,
		tracer:             tracer,
		closer:             closer,
	}, nil
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.closer
}

func (s *Server) CreateProduct(ctx context.Context, product *protopb.Product) (*protopb.EmptyResponseAgent, error) {
	return s.productController.CreateProduct(ctx, product)
}

func (s *Server) GetAllProductsByAgentId(ctx context.Context, agent *protopb.UserAgentApp) (*protopb.ProductsArray, error) {
	return s.productController.GetAllProductsByAgentId(ctx, agent)
}

func (s *Server) GetAllProducts(ctx context.Context, in *protopb.EmptyRequestAgent) (*protopb.ProductsArray, error) {
	return s.productController.GetAllProducts(ctx, in)
}

func (s *Server) GetProductById(ctx context.Context, in *protopb.Product) (*protopb.Product, error) {
	return s.productController.GetProductById(ctx, in)
}

func (s *Server) DeleteProduct(ctx context.Context, in *protopb.Product) (*protopb.EmptyResponseAgent, error) {
	return s.productController.DeleteProduct(ctx, in)
}

func (s *Server) UpdateProduct(ctx context.Context, in *protopb.Product) (*protopb.EmptyResponseAgent, error) {
	return s.productController.UpdateProduct(ctx, in)
}

func (s *Server) OrderProduct(ctx context.Context, in *protopb.Order) (*protopb.EmptyResponseAgent, error) {
	return s.productController.OrderProduct(ctx, in)
}

func (s *Server) GetOrdersByUser(ctx context.Context, in *protopb.UserAgentApp) (*protopb.OrdersArray, error) {
	return s.productController.GetOrdersByUser(ctx, in)
}

func (s *Server) GetOrdersByAgent(ctx context.Context, in *protopb.UserAgentApp) (*protopb.OrdersArray, error) {
	return s.productController.GetOrdersByAgent(ctx, in)
}

func (s *Server) LoginUserInAgentApp(ctx context.Context, login *protopb.LoginRequestAgentApp) (*protopb.LoginResponseAgentApp, error) {
	return s.userController.LoginUserInAgentApp(ctx, login)
}

func (s *Server) CreateUserInAgentApp(ctx context.Context, user *protopb.CreateUserRequestAgentApp) (*protopb.EmptyResponseAgent, error) {
	return s.userController.CreateUserInAgentApp(ctx, user)
}

func (s *Server) GetUserByUsername(ctx context.Context, in *protopb.RequestUsernameAgent) (*protopb.UserAgentApp, error) {
	return s.userController.GetUserByUsername(ctx, in)
}

func (s *Server) GetKeyByUserId(ctx context.Context, in *protopb.RequestIdAgent) (*protopb.ApiTokenAgent, error) {
	return s.userController.GetKeyByUserId(ctx, in)
}

func (s *Server) UpdateKey(ctx context.Context, in *protopb.ApiTokenAgent) (*protopb.EmptyResponseAgent, error) {
	return s.userController.UpdateKey(ctx, in)
}

func (s *Server) CreateCampaignReport(ctx context.Context, in *protopb.RequestIdAgent) (*protopb.EmptyResponseAgent, error) {
	return s.campaignController.CreateCampaignReport(ctx, in)
}
