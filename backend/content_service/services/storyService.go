package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type StoryService struct {
	storyRepository repositories.StoryRepository
	mediaRepository   repositories.MediaRepository
	tagRepository	  repositories.TagRepository
}

func NewStoryService(db *gorm.DB) (*StoryService, error){
	storyRepository, err := repositories.NewStoryRepo(db)
	if err != nil {
		return nil, err
	}

	mediaRepository, err := repositories.NewMediaRepo(db)
	if err != nil {
		return nil, err
	}

	tagRepository, err := repositories.NewTagRepo(db)
	if err != nil {
		return nil, err
	}

	return &StoryService{
		storyRepository,
		mediaRepository,
		tagRepository,
	}, err
}


func (service *StoryService) GetAllHomeStories(ctx context.Context) ([]domain.Story, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHomeStories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbStories, err := service.storyRepository.GetAllHomeStories(ctx)
	if err != nil { return []domain.Story{}, err }
	
	stories := []domain.Story{}
	for _, dbStory := range dbStories{
		story, err := service.retrieveStoryAdditionalData(ctx, dbStory)
		if err != nil { return []domain.Story{}, err }
		stories = append(stories, story)
	}

	return stories, nil
}
func (service *StoryService) CreateStory(ctx context.Context, story *domain.Story) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if len(story.Media) == 0 {
		return errors.New("cannot create empty story")
	}

	return service.storyRepository.CreateStory(ctx, story)
}

func (service *StoryService) GetStoryById(ctx context.Context, id string) (domain.Story, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetStoryById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbStory, err := service.storyRepository.GetStoryById(ctx, id)
	if err != nil { return domain.Story{}, err }

	story, err := service.retrieveStoryAdditionalData(ctx, *dbStory)
	if err != nil { return domain.Story{}, err }

	return story, nil
}

func (service *StoryService) RemoveStory(ctx context.Context, id string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if id == "" {
		return errors.New("cannot remove story")
	}

	return service.storyRepository.RemoveStory(ctx, id)
}

func (service *StoryService) retrieveStoryAdditionalData(ctx context.Context, story persistence.Story) (domain.Story, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "retrieveStoryAdditionalData")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbMedia, err := service.mediaRepository.GetMediaForPost(ctx, story.Id)
	if err != nil { return domain.Story{}, err }
	media := []domain.Media{}
	for _, single := range dbMedia{
		tags, err := service.tagRepository.GetTagsForMedia(ctx, single.Id)
		if err != nil { return domain.Story{}, err }

		converted, err := single.ConvertToDomain(tags)
		if err != nil { return domain.Story{}, err }

		media = append(media, converted)
	}

	converted := story.ConvertToDomain(media)

	return converted, nil
}

func (service *StoryService) GetStoriesFromHighlight(ctx context.Context, highlightId string) ([]domain.Story, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetStoriesFromHighlight")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbStories, err := service.storyRepository.GetHighlightsStories(ctx, highlightId)
	if err != nil { return []domain.Story{}, err }

	stories := []domain.Story{}
	for _, dbStory := range dbStories {
		story, err := service.retrieveStoryAdditionalData(ctx, dbStory)
		if err != nil { return []domain.Story{}, err }
		stories = append(stories, story)
	}

	return stories, nil
}
