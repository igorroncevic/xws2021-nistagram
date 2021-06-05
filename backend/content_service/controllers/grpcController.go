package controllers

import (
	"context"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	otgo "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"io"
)

type Server struct {
	protopb.UnimplementedContentServer
	postController      *PostGrpcController
	commentController   *CommentGrpcController
	likeController      *LikeGrpcController
	favoritesController *FavoritesGrpcController
	hashtagController   *HashtagGrpcController
	tracer              otgo.Tracer
	closer              io.Closer
}

func NewServer(db *gorm.DB) (*Server, error) {
	postController, _ := NewPostController(db)
	commentController, _ := NewCommentController(db)
	likeController, _ := NewLikeController(db)
	favoritesController, _ := NewFavoritesController(db)
	hashtagController, _ := NewHashtagController(db)
	tracer, closer := tracer.Init("global_ContentGrpcController")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		postController:      postController,
		commentController:   commentController,
		likeController:      likeController,
		favoritesController: favoritesController,
		hashtagController:   hashtagController,
		tracer:              tracer,
		closer:              closer,
	}, nil
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.closer
}

/*   Posts   */
func (s *Server) CreatePost(ctx context.Context, in *protopb.Post) (*protopb.EmptyResponseContent, error) {
	return s.postController.CreatePost(ctx, in)
}

func (s *Server) GetAllPosts(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.ReducedPostArray, error) {
	return s.postController.GetAllPosts(ctx, in)
}

func (s *Server) GetPostById(ctx context.Context, in *protopb.RequestId) (*protopb.Post, error) {
	return s.postController.GetPostById(ctx, in.Id)
}

func (s *Server) RemovePost(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	return s.postController.RemovePost(ctx, in.Id)
}

func (s *Server) GetPostsByHashtag(ctx context.Context, in *protopb.Hashtag) (*protopb.ReducedPostArray, error) {
	return s.postController.GetPostsByHashtag(ctx, in)
}

/*   Comments   */
func (s *Server) CreateComment(ctx context.Context, in *protopb.Comment) (*protopb.EmptyResponseContent, error) {
	return s.commentController.CreateComment(ctx, in)
}

func (s *Server) GetCommentsForPost(ctx context.Context, in *protopb.RequestId) (*protopb.CommentsArray, error) {
	return s.commentController.GetCommentsForPost(ctx, in.Id)
}

/* Likes & Dislikes */
func (s *Server) CreateLike(ctx context.Context, in *protopb.Like) (*protopb.EmptyResponseContent, error) {
	return s.likeController.CreateLike(ctx, in)
}

func (s *Server) GetLikesForPost(ctx context.Context, in *protopb.RequestId) (*protopb.LikesArray, error) {
	return s.likeController.GetLikesForPost(ctx, in.Id, true)
}

/* Collections & Favorites */
func (s *Server) GetAllCollections(ctx context.Context, in *protopb.RequestId) (*protopb.CollectionsArray, error) {
	return s.favoritesController.GetAllCollections(ctx, in)
}

func (s *Server) GetCollection(ctx context.Context, in *protopb.RequestId) (*protopb.Collection, error) {
	return s.favoritesController.GetCollection(ctx, in)
}

func (s *Server) CreateCollection(ctx context.Context, in *protopb.Collection) (*protopb.EmptyResponseContent, error) {
	return s.favoritesController.CreateCollection(ctx, in)
}

func (s *Server) RemoveCollection(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	return s.favoritesController.RemoveCollection(ctx, in)
}

func (s *Server) GetUserFavorites(ctx context.Context, in *protopb.RequestId) (*protopb.Favorites, error) {
	return s.favoritesController.GetUserFavorites(ctx, in)
}

func (s *Server) CreateFavorite(ctx context.Context, in *protopb.FavoritesRequest) (*protopb.EmptyResponseContent, error) {
	return s.favoritesController.CreateFavorite(ctx, in)
}

func (s *Server) RemoveFavorite(ctx context.Context, in *protopb.FavoritesRequest) (*protopb.EmptyResponseContent, error) {
	return s.favoritesController.RemoveFavorite(ctx, in)
}

func (s *Server) SearchContentByLocation(ctx context.Context, in *protopb.SearchLocationRequest) (*protopb.ReducedPostArray, error) {
	return s.postController.SearchContentByLocation(ctx, in)
}

/*   Hashtags   */
func (s *Server) CreateHashtag(ctx context.Context, in *protopb.Hashtag) (*protopb.Hashtag, error) {
	return s.hashtagController.CreateHashtag(ctx, in)
}
