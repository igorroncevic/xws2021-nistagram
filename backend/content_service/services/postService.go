package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type PostService struct {
	postRepository repositories.PostRepository
	commentRepository repositories.CommentRepository
	likeRepository 	  repositories.LikeRepository
	mediaRepository   repositories.MediaRepository
	tagRepository	  repositories.TagRepository
}

func NewPostService(db *gorm.DB) (*PostService, error){
	postRepository, err := repositories.NewPostRepo(db)
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

	return &PostService{
		postRepository,
		commentRepository,
		likeRepository,
		mediaRepository,
		tagRepository,
	}, err
}

func (service *PostService) GetAllPosts(ctx context.Context) ([]domain.ReducedPost, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts := []domain.ReducedPost{}

	dbPosts, err := service.postRepository.GetAllPosts(ctx)
	if err != nil{
		return posts, err
	}

	// TODO Retrieve all domain data
	for _, post := range dbPosts{
		converted, err := service.GetReducedPostData(ctx, post.Id)
		if err != nil {
			return []domain.ReducedPost{}, err
		}

		posts = append(posts, converted)
	}

	return posts, nil
}

func (service *PostService) CreatePost(ctx context.Context, post *domain.Post) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if len(post.Media) == 0 {
		return errors.New("cannot create empty post")
	}

	return service.postRepository.CreatePost(ctx, post)
}

func (service *PostService) RemovePost(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemovePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if id == "" {
		return errors.New("cannot remove post")
	}

	return service.postRepository.RemovePost(ctx, id)
}

func (service *PostService) GetPostById(ctx context.Context, id string) (domain.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if id == "" { return domain.Post{}, errors.New("cannot retrieve post") }

	dbPost, err := service.postRepository.GetPostById(ctx, id)
	if err != nil { return domain.Post{}, err }

	dbComments, err := service.commentRepository.GetCommentsForPost(ctx, dbPost.Id)
	if err != nil { return domain.Post{}, err }
	comments := []domain.Comment{}
	for _, comment := range dbComments{
		comments = append(comments, comment.ConvertToDomain("someusername")) // TODO Retrieve usernames from other service
	}

	dbLikes, err := service.likeRepository.GetLikesForPost(ctx, dbPost.Id, true)
	if err != nil { return domain.Post{}, err }
	likes := []domain.Like{}
	for _, like := range dbLikes{
		likes = append(likes, like.ConvertToDomain()) // TODO Retrieve usernames from other service
	}

	dbDislikes, err := service.likeRepository.GetLikesForPost(ctx, dbPost.Id, false)
	if err != nil { return domain.Post{}, err }
	dislikes := []domain.Like{}
	for _, dislike := range dbDislikes{
		dislikes = append(dislikes, dislike.ConvertToDomain()) // TODO Retrieve usernames from other service
	}

	dbMedia, err := service.mediaRepository.GetMediaForPost(ctx, dbPost.Id)
	if err != nil { return domain.Post{}, err }
	media := []domain.Media{}
	for _, single := range dbMedia{
		tags, err := service.tagRepository.GetTagsForMedia(ctx, single.Id)
		if err != nil { return domain.Post{}, err }

		converted, err := single.ConvertToDomain(tags)
		if err != nil { return domain.Post{}, err }

		media = append(media, converted)
	}

	post := dbPost.ConvertToDomain(comments, likes, dislikes, media)

	return post, nil
}

func (service *PostService) GetReducedPostData(ctx context.Context, postId string) (domain.ReducedPost, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetReducedPostData")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	commentsNum, err := service.commentRepository.GetCommentsNumForPost(ctx, postId)
	if err != nil{
		return domain.ReducedPost{}, errors.New("unable to retrieve posts comments")
	}

	likes, err := service.likeRepository.GetLikesNumForPost(ctx, postId, true)
	if err != nil {
		return domain.ReducedPost{}, errors.New("unable to retrieve posts likes")
	}

	dislikes, err := service.likeRepository.GetLikesNumForPost(ctx, postId, false)
	if err != nil {
		return domain.ReducedPost{}, errors.New("unable to retrieve posts dislikes")
	}

	media, err := service.mediaRepository.GetMediaForPost(ctx, postId)
	if err != nil {
		return domain.ReducedPost{}, errors.New("unable to retrieve posts media")
	}

	convertedMedia := []domain.Media{}
	for _, single := range media{
		tags, err := service.tagRepository.GetTagsForMedia(ctx, single.Id)
		if err != nil {
			return domain.ReducedPost{}, errors.New("unable to retrieve media tags")
		}

		converted, err := single.ConvertToDomain(tags)
		if err != nil {
			return domain.ReducedPost{}, errors.New("unable to convert media")
		}

		convertedMedia = append(convertedMedia, converted)
	}

	if err != nil {
		return domain.ReducedPost{}, errors.New("unable to convert posts media")
	}

	post, err := service.postRepository.GetPostById(ctx, postId)
	if err != nil { return domain.ReducedPost{}, err }

	retVal := post.ConvertToDomainReduced(commentsNum, likes, dislikes, convertedMedia)
	return retVal, nil
}
