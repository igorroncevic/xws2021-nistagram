package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type HighlightService struct {
	highlightRepository repositories.HighlightRepository
	storyService        *StoryService
}

func NewHighlightService(db *gorm.DB) (*HighlightService, error){
	highlightRepository, err := repositories.NewHighlightRepo(db)
	if err != nil {
		return nil, err
	}

	storyService, err := NewStoryService(db)
	if err != nil { return nil, err }

	return &HighlightService{
		highlightRepository,
		storyService,
	}, err
}

func (service *HighlightService) GetAllHighlights(ctx context.Context, userId string) ([]domain.Highlight, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHighlights")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbHighlights, err := service.highlightRepository.GetAllHighlights(ctx, userId)
	if err != nil {
		return []domain.Highlight{}, err
	}

	highlights := []domain.Highlight{}
	for _, dbHighlight := range dbHighlights {
		stories, err := service.storyService.GetStoriesFromHighlight(ctx, dbHighlight.Id)
		if err != nil {
			return []domain.Highlight{}, err
		}

		highlight := dbHighlight.ConvertToDomain(stories)
		highlights = append(highlights, highlight)
	}

	return highlights, nil
}

func (service *HighlightService) GetHighlight(ctx context.Context, highlightId string) (domain.Highlight, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetHighlight")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbHighlight, err := service.highlightRepository.GetHighlight(ctx, highlightId)
	if err != nil {
		return domain.Highlight{}, err
	}

	stories, err := service.storyService.GetStoriesFromHighlight(ctx, dbHighlight.Id)
	if err != nil {
		return domain.Highlight{}, err
	}

	highlight := dbHighlight.ConvertToDomain(stories)

	return highlight, nil
}

func (service *HighlightService) CreateHighlightStory(ctx context.Context, highlightRequest domain.HighlightRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHighlightStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.highlightRepository.CreateHighlightStory(ctx, highlightRequest)
	if err != nil{ return err }

	return  nil
}
func (service *HighlightService) RemoveHighlightStory(ctx context.Context, highlightRequest domain.HighlightRequest) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveHighlightStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.highlightRepository.RemoveHighlightStory(ctx, highlightRequest)
	if err != nil { return err }

	return nil
}

func (service *HighlightService) CreateHighlight(ctx context.Context, highlight domain.Highlight) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHighlight")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.highlightRepository.CreateHighlight(ctx, highlight)
	if err != nil {
		return err
	}

	return nil
}
func (service *HighlightService) RemoveHighlight(ctx context.Context, highlightId string, userId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveHighlight")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.highlightRepository.RemoveHighlight(ctx, highlightId, userId)
	if err != nil {
		return err
	}

	return nil
}