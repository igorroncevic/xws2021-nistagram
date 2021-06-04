package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	userspb "github.com/david-drvar/xws2021-nistagram/content_service/proto_intercommunication"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
)

type ContentService struct {
	contentRepository repositories.ContentRepository
	commentRepository repositories.CommentRepository
	likeRepository    repositories.LikeRepository
	mediaRepository   repositories.MediaRepository
	tagRepository     repositories.TagRepository
}

func NewContentService(db *gorm.DB) (*ContentService, error) {
	contentRepository, err := repositories.NewContentRepo(db)
	if err != nil {
		return nil, err
	}

	commentRepository, err := repositories.NewCommentRepo(db)
	if err != nil {
		return nil, err
	}

	likeRepository, err := repositories.NewLikeRepo(db)
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

	return &ContentService{
		contentRepository,
		commentRepository,
		likeRepository,
		mediaRepository,
		tagRepository,
	}, err
}

func (service *ContentService) GetAllPosts(ctx context.Context) ([]domain.ReducedPost, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []domain.ReducedPost{}

	dbPosts, err := service.contentRepository.GetAllPosts(ctx)
	if err != nil {
		return posts, err
	}

	// TODO Retrieve all domain data
	for _, post := range dbPosts {
		commentsNum, err := service.commentRepository.GetCommentsNumForPost(ctx, post.Id)
		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts comments")
		}

		likes, err := service.likeRepository.GetLikesNumForPost(ctx, post.Id, true)
		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts likes")
		}

		dislikes, err := service.likeRepository.GetLikesNumForPost(ctx, post.Id, false)
		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts dislikes")
		}

		media, err := service.mediaRepository.GetMediaForPost(ctx, post.Id)
		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts media")
		}

		convertedMedia := []domain.Media{}
		for _, single := range media {
			tags, err := service.tagRepository.GetTagsForMedia(ctx, single.Id)
			if err != nil {
				return []domain.ReducedPost{}, errors.New("unable to retrieve media tags")
			}

			converted, err := single.ConvertToDomain(tags)
			if err != nil {
				return []domain.ReducedPost{}, errors.New("unable to convert media")
			}

			convertedMedia = append(convertedMedia, converted)
		}

		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to convert posts media")
		}

		posts = append(posts, post.ConvertToDomainReduced(commentsNum, likes, dislikes, convertedMedia))
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

func (service *ContentService) SearchContentByLocation(ctx context.Context, location string) ([]domain.ReducedPost, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []domain.ReducedPost{}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := userspb.NewPrivacyClient(conn)

	dbPosts, err := service.contentRepository.GetPostsByLocation(ctx, location)
	if err != nil {
		return posts, err
	}

	for _, post := range dbPosts {
		commentsNum, err := service.commentRepository.GetCommentsNumForPost(ctx, post.Id)
		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts comments")
		}

		likes, err := service.likeRepository.GetLikesNumForPost(ctx, post.Id, true)
		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts likes")
		}

		dislikes, err := service.likeRepository.GetLikesNumForPost(ctx, post.Id, false)
		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts dislikes")
		}

		media, err := service.mediaRepository.GetMediaForPost(ctx, post.Id)
		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to retrieve posts media")
		}

		convertedMedia := []domain.Media{}
		for _, single := range media {
			tags, err := service.tagRepository.GetTagsForMedia(ctx, single.Id)
			if err != nil {
				return []domain.ReducedPost{}, errors.New("unable to retrieve media tags")
			}

			converted, err := single.ConvertToDomain(tags)
			if err != nil {
				return []domain.ReducedPost{}, errors.New("unable to convert media")
			}

			convertedMedia = append(convertedMedia, converted)
		}

		if err != nil {
			return []domain.ReducedPost{}, errors.New("unable to convert posts media")
		}

		posts = append(posts, post.ConvertToDomainReduced(commentsNum, likes, dislikes, convertedMedia))
	}

	finalPosts := []domain.ReducedPost{}

	//call user service to check if users profile is public
	for _, post := range posts {
		privacyRequest := userspb.PrivacyRequest{
			UserId: post.UserId,
		}
		response, err := c.CheckUserProfilePublic(context.Background(), &privacyRequest)
		if err != nil {
			log.Fatalf("Error when calling CheckUserProfilePublic: %s", err)
		}
		if response.Response {
			finalPosts = append(finalPosts, post)
		}
	}

	return finalPosts, nil
}
