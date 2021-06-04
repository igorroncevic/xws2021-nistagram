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

type HashtagGrpcController struct {
	service *services.HashtagService
}

func NewHashtagController(db *gorm.DB) (*HashtagGrpcController, error) {
	service, err := services.NewHashtagService(db)
	if err != nil {
		return nil, err
	}

	return &HashtagGrpcController{
		service,
	}, nil
}

func (s *HashtagGrpcController) CreateHashtag(ctx context.Context, in *contentpb.Hashtag) (*contentpb.Hashtag, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var hashtag *domain.Hashtag
	hashtag = hashtag.ConvertFromGrpc(in)

	hashtag, err := s.service.CreateHashtag(ctx, hashtag.Text)
	if err != nil {
		return &contentpb.Hashtag{}, status.Errorf(codes.Unknown, "could not create hashtag")
	}

	return &contentpb.Hashtag{Id: hashtag.Id, Text: hashtag.Text}, nil
}
