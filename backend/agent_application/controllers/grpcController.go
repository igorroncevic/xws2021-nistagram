package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	otgo "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"io"
)

type Server struct {
	protopb.UnimplementedAgentServer
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger) (*Server, error) {
	tracer, closer := tracer.Init("agentService")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		tracer: tracer,
		closer: closer,
	}, nil
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.closer
}

func (s *Server) CreateProduct(ctx context.Context, product *protopb.Product) (*protopb.EmptyResponseAgent, error) {
	panic("implement me")
}
