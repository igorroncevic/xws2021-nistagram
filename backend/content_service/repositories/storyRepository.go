package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/content_service/util/images"
	"gorm.io/gorm"
)

type StoryRepository interface {
	GetAllHomeStories(context.Context, []string, bool) ([]persistence.Story, error)
	CreateStory(context.Context, *domain.Story) error
	GetStoryById(context.Context, string) (*persistence.Story, error)
	RemoveStory(context.Context, string, string) error
	GetHighlightsStories(context.Context, string) ([]persistence.Story, error)
	GetUsersStories(context.Context, string, bool) ([]persistence.Story, error)
}

type storyRepository struct {
	DB *gorm.DB
	mediaRepository MediaRepository
	tagRepository   TagRepository
}

func NewStoryRepo(db *gorm.DB) (*storyRepository, error) {
	if db == nil {
		panic("StoryRepository not created, gorm.DB is nil")
	}

	mediaRepository, _ := NewMediaRepo(db)
	tagRepository, _ := NewTagRepo(db)

	return &storyRepository{
		DB: db,
		mediaRepository: mediaRepository,
		tagRepository: tagRepository,
	}, nil
}

const(
	storyDurationQuery = "created_at > (LOCALTIMESTAMP - interval '1 day')"
)

func (repository *storyRepository) GetAllHomeStories(ctx context.Context, userIds []string, isCloseFriends bool) ([]persistence.Story, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllHomeStories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	stories := []persistence.Story{}
	result := repository.DB.Order("created_at desc").
		Where(storyDurationQuery + " AND user_id IN ? AND is_close_friends = ?", userIds, isCloseFriends).
		Find(&stories)
	if result.Error != nil {
		return stories, result.Error
	}

	return stories, nil
}

func (repository *storyRepository) GetUsersStories(ctx context.Context, userId string, isCloseFriends bool) ([]persistence.Story, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsersStories")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	// TODO Retrieve ids from people you follow
	stories := []persistence.Story{}
	result := repository.DB.Order("created_at desc").
		Where(storyDurationQuery + " AND user_id = ? AND is_close_friends = ?", userId, isCloseFriends).
		Find(&stories)
	if result.Error != nil {
		return stories, result.Error
	}

	return stories, nil
}

// Internal use only
func (repository *storyRepository) GetStoryById(ctx context.Context, id string) (*persistence.Story, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetStoryById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	story := &persistence.Story{}
	result := repository.DB.Where("id = ? AND " + storyDurationQuery, id).First(&story)
	if result.Error != nil {
		return story, result.Error
	}

	return story, nil
}


func (repository *storyRepository) CreateStory(ctx context.Context, story *domain.Story) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func (tx *gorm.DB) error {
		var storyToSave persistence.Story
		storyToSave = storyToSave.ConvertToPersistence(*story)

		result := repository.DB.Create(&storyToSave)
		if result.Error != nil || result.RowsAffected != 1 {
			return errors.New("cannot save story")
		}

		for _, media := range story.Media{
			media.PostId = storyToSave.Id
			dbMedia, err := repository.mediaRepository.CreateMedia(ctx, media)

			if err != nil{
				return errors.New("cannot save story")
			}

			for _, tag := range media.Tags{
				tag.MediaId = dbMedia.Id
				err := repository.tagRepository.CreateTag(ctx, tag)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil { return err }
	return nil
}

func (repository *storyRepository) RemoveStory(ctx context.Context, storyId string, userId string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := repository.DB.Transaction(func (tx *gorm.DB) error{
		story := &persistence.Story{
			Post: persistence.Post{Id: storyId, UserId: userId },
		}
		result := repository.DB.First(&story)

		if result.Error != nil || result.RowsAffected != 1 {
			return errors.New("cannot remove non-existing story")
		}

		result = repository.DB.Delete(&story)
		if result.Error != nil || result.RowsAffected != 1 {
			return errors.New("cannot remove story")
		}

		storyMedia, err := repository.mediaRepository.GetMediaForPost(ctx, story.Id)
		if err != nil {
			return errors.New("cannot retrieve story's media")
		}

		for _, media := range storyMedia {
			err := tx.Transaction(func (tx *gorm.DB) error {
				mediaTags, err := repository.tagRepository.GetTagsForMedia(ctx, media.Id)
				if err != nil { return err }

				for _, tag := range mediaTags{
					var tagPers *persistence.Tag
					tagPers = tagPers.ConvertToPersistence(tag)

					err := repository.tagRepository.RemoveTag(ctx, *tagPers)
					if err != nil { return err }
				}

				err = images.RemoveImages([]string{media.Filename})
				if err != nil { return errors.New("cannot remove media's images") }

				err = repository.mediaRepository.RemoveMedia(ctx, media.Id)
				if err != nil{ return errors.New("cannot remove story's media") }

				return nil
			})
			if err != nil { return err }
		}
		return nil
	})

	if err != nil { return err }
	return nil
}

func (repository *storyRepository) GetHighlightsStories (ctx context.Context, highlightId string) ([]persistence.Story, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveStory")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	stories := []persistence.Story{}
	result := repository.DB.Model(&persistence.Story{}).
		Joins("left join highlight_stories ON stories.id = highlight_stories.story_id").
		Joins("left join highlights 		 ON highlight_stories.highlight_id = highlights.id").
		Where("highlights.id = ?", highlightId).
		Find(&stories)

	if result.Error != nil {
		return stories, result.Error
	}

	return stories, nil
}