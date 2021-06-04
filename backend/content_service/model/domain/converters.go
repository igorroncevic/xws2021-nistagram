package domain

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model"
	contentpb "github.com/david-drvar/xws2021-nistagram/content_service/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Post Converters
func (p *Post) ConvertFromGrpc(post *contentpb.Post) *Post {
	if post == nil {
		post = &contentpb.Post{}
	}
	return &Post{
		Objava: Objava{
			Id:          post.Id,
			UserId:      post.UserId,
			IsAd:        post.IsAd,
			Type:        model.GetPostType(post.Type),
			Description: post.Description,
			Location:    post.Location,
			CreatedAt:   post.CreatedAt.AsTime(),
			Media:       ConvertMultipleMediaFromGrpc(post.Media),
		},
		Comments: ConvertMultipleCommentsFromGrpc(post.Comments),
		Likes:    ConvertMultipleLikesFromGrpc(post.Likes),
		Dislikes: ConvertMultipleLikesFromGrpc(post.Dislikes),
	}
}

func (p Post) ConvertToGrpc() *contentpb.Post {
	return &contentpb.Post{
		Id:          p.Id,
		UserId:      p.UserId,
		IsAd:        p.IsAd,
		Type:        p.Type.String(),
		Description: p.Description,
		Location:    p.Location,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		Media:       ConvertMultipleMediaToGrpc(p.Media),
		Comments:    ConvertMultipleCommentsToGrpc(p.Comments),
		Likes:       ConvertMultipleLikesToGrpc(p.Likes),
		Dislikes:    ConvertMultipleLikesToGrpc(p.Dislikes),
	}
}

func ConvertMultiplePostsToGrpc(posts []Post) []*contentpb.Post {
	grpcPosts := []*contentpb.Post{}
	for _, post := range posts {
		grpcPosts = append(grpcPosts, post.ConvertToGrpc())
	}

	return grpcPosts
}

// ReducedPost Converters
func (p ReducedPost) ConvertToGrpc() *contentpb.ReducedPost {
	return &contentpb.ReducedPost{
		Id:          p.Id,
		UserId:      p.UserId,
		IsAd:        p.IsAd,
		Type:        p.Type.String(),
		Description: p.Description,
		Location:    p.Location,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		Media:       ConvertMultipleMediaToGrpc(p.Media),
		CommentsNum: p.CommentsNum,
		LikesNum:    p.LikesNum,
		DislikesNum: p.DislikesNum,
	}
}

func (p ReducedPost) ConvertFromGrpc(post *contentpb.ReducedPost) ReducedPost {
	return ReducedPost{
		Objava: Objava{
			Id:          post.Id,
			UserId:      post.UserId,
			IsAd:        post.IsAd,
			Type:        model.GetPostType(post.Type),
			Description: post.Description,
			Location:    post.Location,
			CreatedAt:   post.CreatedAt.AsTime(),
			Media:       ConvertMultipleMediaFromGrpc(post.Media),
		},
		CommentsNum: post.CommentsNum,
		LikesNum:    post.LikesNum,
		DislikesNum: post.DislikesNum,
	}
}

func ConvertMultipleReducedPostsToGrpc(posts []ReducedPost) []*contentpb.ReducedPost {
	grpcPosts := []*contentpb.ReducedPost{}
	for _, post := range posts {
		grpcPosts = append(grpcPosts, post.ConvertToGrpc())
	}

	return grpcPosts
}

func ConvertMultipleReducedPostsFromGrpc(posts []*contentpb.ReducedPost) []ReducedPost {
	convertedPosts := []ReducedPost{}
	for _, post := range posts {
		var converted ReducedPost
		converted = converted.ConvertFromGrpc(post)

		convertedPosts = append(convertedPosts, converted)
	}

	return convertedPosts
}

// Media converters
func (m *Media) ConvertFromGrpc(media *contentpb.Media) *Media {
	if m == nil {
		m = &Media{}
	}
	return &Media{
		Id:       media.Id,
		Type:     model.GetMediaType(media.Type),
		PostId:   media.PostId,
		Content:  media.Content,
		OrderNum: media.OrderNum,
		Tags:     ConvertMultipleTagFromGrpc(media.Tags),
	}
}

func (m *Media) ConvertToGrpc() *contentpb.Media {
	if m == nil {
		m = &Media{}
	}
	return &contentpb.Media{
		Id:       m.Id,
		Type:     m.Type.String(),
		PostId:   m.PostId,
		Content:  m.Content,
		OrderNum: m.OrderNum,
		Tags:     ConvertMultipleTagsToGrpc(m.Tags),
	}
}

func ConvertMultipleMediaFromGrpc(m []*contentpb.Media) []Media {
	media := []Media{}
	if m != nil {
		for _, protoMedia := range m {
			var domainMedia *Media
			domainMedia = domainMedia.ConvertFromGrpc(protoMedia)
			media = append(media, *domainMedia)
		}
	}
	return media
}

func ConvertMultipleMediaToGrpc(m []Media) []*contentpb.Media {
	media := []*contentpb.Media{}
	if m != nil {
		for _, domainMedia := range m {
			protoMedia := domainMedia.ConvertToGrpc()
			media = append(media, protoMedia)
		}
	}
	return media
}

// Comment converters
func (c *Comment) ConvertFromGrpc(comment *contentpb.Comment) *Comment {
	if c == nil {
		c = &Comment{}
	}
	return &Comment{
		Id:        comment.Id,
		PostId:    comment.PostId,
		UserId:    comment.UserId,
		Username:  comment.Username,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.AsTime(),
	}
}

func (c *Comment) ConvertToGrpc() *contentpb.Comment {
	if c == nil {
		c = &Comment{}
	}
	return &contentpb.Comment{
		Id:        c.Id,
		PostId:    c.PostId,
		UserId:    c.UserId,
		Username:  c.Username,
		Content:   c.Content,
		CreatedAt: timestamppb.New(c.CreatedAt),
	}
}

func ConvertMultipleCommentsFromGrpc(c []*contentpb.Comment) []Comment {
	comments := []Comment{}
	if c != nil {
		for _, protoComment := range c {
			var domainComment *Comment
			domainComment = domainComment.ConvertFromGrpc(protoComment)
			comments = append(comments, *domainComment)
		}
	}
	return comments
}

func ConvertMultipleCommentsToGrpc(c []Comment) []*contentpb.Comment {
	comments := []*contentpb.Comment{}
	if c != nil {
		for _, domainComment := range c {
			protoComment := domainComment.ConvertToGrpc()
			comments = append(comments, protoComment)
		}
	}
	return comments
}

// Tag converters
func (t *Tag) ConvertFromGrpc(tag *contentpb.Tag) *Tag {
	if t == nil {
		t = &Tag{}
	}
	return &Tag{
		MediaId:  tag.MediaId,
		UserId:   tag.UserId,
		Username: tag.Username,
	}
}

func (t *Tag) ConvertToGrpc() *contentpb.Tag {
	if t == nil {
		t = &Tag{}
	}
	return &contentpb.Tag{
		MediaId:  t.MediaId,
		UserId:   t.UserId,
		Username: t.Username,
	}
}

func ConvertMultipleTagFromGrpc(t []*contentpb.Tag) []Tag {
	tags := []Tag{}
	if t != nil {
		for _, protoTag := range t {
			var domainTag *Tag
			domainTag = domainTag.ConvertFromGrpc(protoTag)
			tags = append(tags, *domainTag)
		}
	}

	return tags
}

func ConvertMultipleTagsToGrpc(t []Tag) []*contentpb.Tag {
	tags := []*contentpb.Tag{}
	if t != nil {
		for _, domainTag := range t {
			tags = append(tags, domainTag.ConvertToGrpc())
		}
	}
	return tags
}

// Like converters
func (l *Like) ConvertFromGrpc(like *contentpb.Like) *Like {
	if l == nil {
		l = &Like{}
	}
	return &Like{
		PostId: like.PostId,
		UserId: like.UserId,
		IsLike: like.IsLike,
	}
}

func (l *Like) ConvertToGrpc() *contentpb.Like {
	if l == nil {
		l = &Like{}
	}
	return &contentpb.Like{
		PostId: l.PostId,
		UserId: l.UserId,
		IsLike: l.IsLike,
	}
}

func ConvertMultipleLikesFromGrpc(l []*contentpb.Like) []Like {
	likes := []Like{}
	if l != nil {
		for _, protoLike := range l {
			var domainLike *Like
			domainLike.ConvertFromGrpc(protoLike)
			likes = append(likes, *domainLike)
		}
	}

	return likes
}

func ConvertMultipleLikesToGrpc(l []Like) []*contentpb.Like {
	likes := []*contentpb.Like{}
	if l != nil {
		for _, domainLike := range l {
			likes = append(likes, domainLike.ConvertToGrpc())
		}
	}
	return likes
}

func (c Collection) ConvertFromGrpc(collection *contentpb.Collection) Collection {
	return Collection{
		Id:     collection.Id,
		Name:   collection.Name,
		UserId: collection.UserId,
		Posts:  ConvertMultipleReducedPostsFromGrpc(collection.Posts),
	}
}

func (c Collection) ConvertToGrpc() *contentpb.Collection {
	return &contentpb.Collection{
		Id:     c.Id,
		Name:   c.Name,
		UserId: c.UserId,
		Posts:  ConvertMultipleReducedPostsToGrpc(c.Posts),
	}
}

func ConvertMultipleCollectionsToGrpc(collections []Collection) []*contentpb.Collection {
	converted := []*contentpb.Collection{}
	for _, collection := range collections {
		converted = append(converted, collection.ConvertToGrpc())
	}

	return converted
}

func (f Favorites) ConvertToGrpc() *contentpb.Favorites {
	return &contentpb.Favorites{
		UserId:       f.UserId,
		Collections:  ConvertMultipleCollectionsToGrpc(f.Collections),
		Unclassified: ConvertMultipleReducedPostsToGrpc(f.Unclassified),
	}
}

func (fr *FavoritesRequest) ConvertFromGrpc(request *contentpb.FavoritesRequest) FavoritesRequest {
	return FavoritesRequest{
		PostId:       request.PostId,
		CollectionId: request.CollectionId,
		UserId:       request.UserId,
	}
}

// Hashtag converters
func (h *Hashtag) ConvertFromGrpc(hashtag *contentpb.Hashtag) *Hashtag {
	if h == nil {
		h = &Hashtag{}
	}
	return &Hashtag{
		Id:   hashtag.Id,
		Text: hashtag.Text,
	}
}
