package controllers

import (
	"context"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
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

func (s *HashtagGrpcController) CreateHashtag(ctx context.Context, in *protopb.Hashtag) (*protopb.Hashtag, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var hashtag *domain.Hashtag
	hashtag = hashtag.ConvertFromGrpc(in)

	hashtag, err := s.service.CreateHashtag(ctx, hashtag.Text)
	if err != nil {
		return &protopb.Hashtag{}, status.Errorf(codes.Unknown, "could not create hashtag")
	}

	return &protopb.Hashtag{Id: hashtag.Id, Text: hashtag.Text}, nil
}

func (s *HashtagGrpcController) GetAllHashtags(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.Hashtags, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHashtags")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	hashtags, err := s.service.GetAllHashtags(ctx)
	if err != nil {
		return &protopb.Hashtags{}, status.Errorf(codes.Unknown, "could not get hashtags")
	}

	responseHashtags := []*protopb.Hashtag{}

	for _, hashtag := range hashtags {
		responseHashtags = append(responseHashtags, &protopb.Hashtag{Id: hashtag.Id, Text: hashtag.Text})
	}

	return &protopb.Hashtags{Hashtags: responseHashtags}, nil
}
