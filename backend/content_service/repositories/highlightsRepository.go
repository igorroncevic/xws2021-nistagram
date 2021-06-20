package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"gorm.io/gorm"
)

type HighlightRepository interface {
	GetAllHighlights(context.Context, string) 			 		 ([]persistence.Highlight, error)
	GetHighlight(context.Context, string) 		 		 		 (persistence.Highlight, error)

	CreateHighlightStory(context.Context, domain.HighlightRequest) error
	RemoveHighlightStory(context.Context, domain.HighlightRequest) error

	CreateHighlight(context.Context, domain.Highlight) (persistence.Highlight, error)
	RemoveHighlight(context.Context, string, string)   error
}

type highlightRepository struct {
	DB                *gorm.DB
	storyRepository StoryRepository
}

func NewHighlightRepo(db *gorm.DB) (*highlightRepository, error) {
	if db == nil {
		panic("HighlightRepository not created, gorm.DB is nil")
	}

	storyRepository, _ := NewStoryRepo(db)

	return &highlightRepository{ DB: db, storyRepository: storyRepository }, nil
}

func (repository *highlightRepository) GetAllHighlights(ctx context.Context, userId string) ([]persistence.Highlight, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHighlights")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	highlights := []persistence.Highlight{}
	result := repository.DB.Where("user_id = ?", userId).Find(&highlights)

	if result.Error != nil {
		return highlights, result.Error
	}

	return highlights, nil
}

func (repository *highlightRepository) GetHighlight(ctx context.Context, id string) (persistence.Highlight, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetHighlight")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	highlight := persistence.Highlight{}
	result := repository.DB.Where("id = ?", id).Find(&highlight)

	if result.Error != nil {
		return highlight, result.Error
	}

	return highlight, nil
}

func (repository *highlightRepository) CreateHighlightStory(ctx context.Context, highlightRequest domain.HighlightRequest) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHighlightStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var highlightPers *persistence.HighlightStory
	highlightPers = highlightPers.ConvertToPersistence(highlightRequest)

	if highlightPers.StoryId != "" {
		// Check if user has that story
		var count int64
		result := repository.DB.Model(&persistence.Story{}).
			Where("id = ? AND user_id = ?", highlightPers.StoryId, highlightRequest.UserId).Count(&count)

		if result.Error != nil {
			return result.Error
		}else if count == 0 {
			return errors.New("user does not own that story")
		}
	}else{
		return errors.New("no story id provided")
	}

	if highlightPers.HighlightId != "" {
		// Check if highlight exists
		var count int64
		result := repository.DB.Model(&persistence.Highlight{}).
			Where("id = ?", highlightPers.HighlightId).Count(&count)

		if result.Error != nil {
			return result.Error
		} else if count == 0 {
			return errors.New("highlight does not exist")
		}
	}else{
		return errors.New("no highlight id provided")
	}

	result := repository.DB.Create(highlightPers)

	if result.Error != nil || result.RowsAffected != 1 {
		return result.Error
	}

	return nil
}

func (repository *highlightRepository) RemoveHighlightStory(ctx context.Context, request domain.HighlightRequest) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveHighlightStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var favoritesPers *persistence.HighlightStory
	favoritesPers = favoritesPers.ConvertToPersistence(request)

	result := repository.DB.Delete(&favoritesPers)

	if result.Error != nil || result.RowsAffected != 1 {
		return result.Error
	}

	return nil
}

func (repository *highlightRepository) CreateHighlight(ctx context.Context, highlight domain.Highlight) (persistence.Highlight, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHighlight")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var highlightPers *persistence.Highlight
	highlightPers = highlightPers.ConvertToPersistence(highlight)
	err := repository.DB.Transaction(func (tx *gorm.DB) error{
		result := repository.DB.Create(highlightPers)

		if result.Error != nil || result.RowsAffected != 1 {
			return result.Error
		}

		// Case: New highlight was created upon saving story to highlights
		if len(highlight.Stories) > 0 {
			for _, story := range highlight.Stories {
				err := repository.CreateHighlightStory(ctx, domain.HighlightRequest{
					UserId:      highlightPers.UserId,
					HighlightId: highlightPers.Id,
					StoryId:     story.Id,
				})
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil { return persistence.Highlight{}, err }
	return *highlightPers, nil
}

func (repository *highlightRepository) RemoveHighlight(ctx context.Context, highlightId string, userId string) error{
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveHighlight")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func (tx *gorm.DB) error{
		highlightPers := persistence.Highlight{Id: highlightId, UserId: userId}
		result := repository.DB.Delete(&highlightPers)

		if result.Error != nil || result.RowsAffected != 1 {
			return result.Error
		}

		highlightStorys, err := repository.storyRepository.GetHighlightsStories(ctx, highlightId)
		if err != nil { return err }

		for _, story := range highlightStorys {
			err := repository.RemoveHighlightStory(ctx, domain.HighlightRequest{
				StoryId:       story.Id,
				HighlightId: highlightId,
			})
			if err != nil { return err }
		}

		return nil
	})

	if err != nil { return err }
	return nil
}

