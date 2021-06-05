package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	contentpb "github.com/david-drvar/xws2021-nistagram/content_service/proto"
	otgo "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"io"
)

type Server struct {
	contentpb.UnimplementedContentServer
	hashtagController   *HashtagGrpcController
	postController 		*PostGrpcController
	commentController 	*CommentGrpcController
	likeController	  	*LikeGrpcController
	favoritesController *FavoritesGrpcController
	storyController		*StoryGrpcController
	highlightController *HighlightGrpcController
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(db *gorm.DB) (*Server, error) {
	postController, _ := NewPostController(db)
	storyController, _ := NewStoryController(db)
	commentController, _ := NewCommentController(db)
	likeController, _ := NewLikeController(db)
	favoritesController, _ := NewFavoritesController(db)
	hashtagController, _ := NewHashtagController(db)
	highlightController, _ := NewHighlightController(db)
	tracer, closer := tracer.Init("global_ContentGrpcController")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		postController:      postController,
		commentController:   commentController,
		likeController:      likeController,
		favoritesController: favoritesController,
		hashtagController:   hashtagController,
		storyController: storyController,
		highlightController: highlightController,
		tracer: tracer,
		closer: closer,
	}, nil
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.closer
}

/*   Posts   */
func (s *Server) CreatePost(ctx context.Context, in *contentpb.Post) (*contentpb.EmptyResponse, error) {
	return s.postController.CreatePost(ctx, in)
}

func (s *Server) GetAllPosts(ctx context.Context, in *contentpb.EmptyRequest) (*contentpb.ReducedPostArray, error) {
	return s.postController.GetAllPosts(ctx, in)
}

func (s *Server) GetPostById(ctx context.Context, in *contentpb.RequestId) (*contentpb.Post, error) {
	return s.postController.GetPostById(ctx, in.Id)
}

func (s *Server) RemovePost(ctx context.Context, in *contentpb.RequestId) (*contentpb.EmptyResponse, error) {
	return s.postController.RemovePost(ctx, in.Id)
}

func (s *Server) GetPostsByHashtag(ctx context.Context, in *contentpb.Hashtag) (*contentpb.ReducedPostArray, error) {
	return s.postController.GetPostsByHashtag(ctx, in)
}

/*   Stories   */
func (s *Server) CreateStory(ctx context.Context, in *contentpb.Story) (*contentpb.EmptyResponse, error) {
	return s.storyController.CreateStory(ctx, in)
}

func (s *Server) GetAllStories(ctx context.Context, in *contentpb.EmptyRequest) (*contentpb.StoriesArray, error) {
	return s.storyController.GetAllHomeStories(ctx, in)
}

func (s *Server) GetStoryById(ctx context.Context, in *contentpb.RequestId) (*contentpb.Story, error) {
	return s.storyController.GetStoryById(ctx, in)
}

func (s *Server) RemoveStory(ctx context.Context, in *contentpb.RequestId) (*contentpb.EmptyResponse, error) {
	return s.storyController.RemoveStory(ctx, in)
}

/*   Comments   */
func (s *Server) CreateComment(ctx context.Context, in *contentpb.Comment) (*contentpb.EmptyResponse, error) {
	return s.commentController.CreateComment(ctx, in)
}

func (s *Server) GetCommentsForPost(ctx context.Context, in *contentpb.RequestId) (*contentpb.CommentsArray, error) {
	return s.commentController.GetCommentsForPost(ctx, in.Id)
}

/* Likes & Dislikes */
func (s *Server) CreateLike(ctx context.Context, in *contentpb.Like) (*contentpb.EmptyResponse, error) {
	return s.likeController.CreateLike(ctx, in)
}

func (s *Server) GetLikesForPost(ctx context.Context, in *contentpb.RequestId) (*contentpb.LikesArray, error) {
	return s.likeController.GetLikesForPost(ctx, in.Id, true)
}

/* Collections & Favorites */
func (s *Server) GetAllCollections(ctx context.Context, in *contentpb.RequestId) (*contentpb.CollectionsArray, error) {
	return s.favoritesController.GetAllCollections(ctx, in)
}

func (s *Server) GetCollection(ctx context.Context, in *contentpb.RequestId) (*contentpb.Collection, error) {
	return s.favoritesController.GetCollection(ctx, in)
}

func (s *Server) CreateCollection(ctx context.Context, in *contentpb.Collection) (*contentpb.EmptyResponse, error) {
	return s.favoritesController.CreateCollection(ctx, in)
}

func (s *Server) RemoveCollection(ctx context.Context, in *contentpb.RequestId) (*contentpb.EmptyResponse, error) {
	return s.favoritesController.RemoveCollection(ctx, in)
}

func (s *Server) GetUserFavorites(ctx context.Context, in *contentpb.RequestId) (*contentpb.Favorites, error) {
	return s.favoritesController.GetUserFavorites(ctx, in)
}

func (s *Server) CreateFavorite(ctx context.Context, in *contentpb.FavoritesRequest) (*contentpb.EmptyResponse, error) {
	return s.favoritesController.CreateFavorite(ctx, in)
}

func (s *Server) RemoveFavorite(ctx context.Context, in *contentpb.FavoritesRequest) (*contentpb.EmptyResponse, error) {
	return s.favoritesController.RemoveFavorite(ctx, in)
}


func (s *Server) SearchContentByLocation(ctx context.Context, in *contentpb.SearchLocationRequest) (*contentpb.ReducedPostArray, error) {
	return s.postController.SearchContentByLocation(ctx, in)
}

/*   Hashtags   */
func (s *Server) CreateHashtag(ctx context.Context, in *contentpb.Hashtag) (*contentpb.Hashtag, error) {
	return s.hashtagController.CreateHashtag(ctx, in)
}

/*   Highlights   */
func (s *Server) GetAllHighlights(ctx context.Context, in *contentpb.RequestId) (*contentpb.HighlightsArray, error) {
	return s.highlightController.GetAllHighlights(ctx, in)
}

func (s *Server) GetHighlight(ctx context.Context, in *contentpb.RequestId) (*contentpb.Highlight, error) {
	return s.highlightController.GetHighlight(ctx, in)
}

func (s *Server) CreateHighlight(ctx context.Context, in *contentpb.Highlight) (*contentpb.EmptyResponse, error) {
	return s.highlightController.CreateHighlight(ctx, in)
}

func (s *Server) RemoveHighlight(ctx context.Context, in *contentpb.RequestId) (*contentpb.EmptyResponse, error) {
	return s.highlightController.RemoveHighlight(ctx, in)
}

func (s *Server) CreateHighlightStory(ctx context.Context, in *contentpb.HighlightRequest) (*contentpb.EmptyResponse, error) {
	return s.highlightController.CreateHighlightStory(ctx, in)
}

func (s *Server) RemoveHighlightStory(ctx context.Context, in *contentpb.HighlightRequest) (*contentpb.EmptyResponse, error) {
	return s.highlightController.RemoveHighlightStory(ctx, in)
}
