package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type ContentService struct {
	contentRepository repositories.ContentRepository
	commentRepository repositories.CommentRepository
}

func NewContentService(db *gorm.DB) (*ContentService, error){
	contentRepository, err := repositories.NewContentRepo(db)
	if err != nil {
		return nil, err
	}

	commentRepository, err := repositories.NewCommentRepo(db)
	if err != nil {
		return nil, err
	}

	return &ContentService{
		contentRepository,
		commentRepository,
	}, err
}

// TODO Use ReducedPost to reduce amount of data being transfered
func (service *ContentService) GetAllPosts(ctx context.Context) ([]domain.ReducedPost, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []domain.ReducedPost{}

	dbPosts, err := service.contentRepository.GetAllPosts(ctx)
	if err != nil{
		return posts, err
	}

	// TODO Retrieve all domain data
	for _, post := range dbPosts{
		commentsNum, err := service.commentRepository.GetCommentsNumForPost(ctx, post.Id)
		if err != nil{
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts comments")
		}
		/*likes := []domain.Like{}
		dislikes := []domain.Like{}*/
		tags := []domain.Tag{}
		media := []domain.Media{}
		posts = append(posts, post.ConvertToDomainReduced(commentsNum, 0, 0, tags, media))
	}

	return posts, nil
}

func (service *ContentService) CreatePost(ctx context.Context, post *domain.Post) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if len(post.Media) == 0 {
		return errors.New("cannot create empty post")
	}

	return service.contentRepository.CreatePost(ctx, post)
}
