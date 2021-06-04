package persistence

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/util"
	"github.com/david-drvar/xws2021-nistagram/content_service/util/images"
	uuid "github.com/satori/go.uuid"
	"time"
)

func (p Post) ConvertToDomain(comments []domain.Comment, likes []domain.Like, dislikes []domain.Like, media []domain.Media) domain.Post{
	return domain.Post{
		Objava:   domain.Objava{
			Id:          p.Id,
			UserId:      p.UserId,
			IsAd:        p.IsAd,
			Type:        p.Type,
			Description: p.Description,
			Location:    p.Location,
			CreatedAt:   p.CreatedAt,
			Media:       media,
		},
		Comments: comments,
		Likes:    likes,
		Dislikes: dislikes,
	}
}

func (p Post) ConvertToDomainReduced(commentsNum int, likesNum int, dislikesNum int, media []domain.Media) domain.ReducedPost{
	return domain.ReducedPost{
		Objava:   domain.Objava{
			Id:          p.Id,
			UserId:      p.UserId,
			IsAd:        p.IsAd,
			Type:        p.Type,
			Description: p.Description,
			Location:    p.Location,
			CreatedAt:   p.CreatedAt,
			Media:       media,
		},
		CommentsNum: 	int32(commentsNum),
		LikesNum:    	int32(likesNum),
		DislikesNum: 	int32(dislikesNum),
	}
}

func (p Post) ConvertToPersistence(post domain.Post) Post {
	return Post{
		Id:          uuid.NewV4().String(),
		UserId:      post.UserId,
		IsAd:        post.IsAd,
		Type:        post.Type,
		Description: post.Description,
		Location:    post.Location,
		CreatedAt:   time.Now(),
	}
}

func (s Story) ConvertToPersistence(story domain.Story) Story {
	return Story{
		Post: Post {
			Id:          uuid.NewV4().String(),
			UserId:      story.UserId,
			IsAd:        story.IsAd,
			Type:        story.Type,
			Description: story.Description,
			Location:    story.Location,
			CreatedAt:   time.Now(),
		},
		IsCloseFriends: story.IsCloseFriends,
	}
}

func (c Comment) ConvertToDomain(username string) domain.Comment {
	return domain.Comment{
		Id:		   c.Id,
		PostId:    c.PostId,
		UserId:    c.UserId,
		Username:  username,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}
}

func (c *Comment) ConvertToPersistence(comment domain.Comment) *Comment{
	if c == nil { c = &Comment{} }
	return &Comment{
		Id:		   uuid.NewV4().String(),
		PostId:    comment.PostId,
		UserId:    comment.UserId,
		Content:   comment.Content,
		CreatedAt: time.Now(),
	}
}

func (l Like) ConvertToDomain() domain.Like {
	return domain.Like{
		PostId: l.PostId,
		UserId: l.UserId,
		IsLike: l.IsLike,
	}
}

func (l *Like) ConvertToPersistence(like domain.Like) *Like{
	if l == nil { l = &Like{} }
	return &Like{
		PostId:    	like.PostId,
		UserId:    	like.UserId,
		IsLike:		like.IsLike,
	}
}

func (m *Media) ConvertToDomain(tags []domain.Tag) (domain.Media, error){
	if m == nil { m = &Media{} }
	filename, err := images.LoadImageToBase64(util.GetContentLocation(m.Filename))
	if err != nil {
		return domain.Media{}, err
	}
	return domain.Media{
		Id:       m.Id,
		Type:     m.Type,
		PostId:   m.PostId,
		Content:  filename,
		OrderNum: int32(m.OrderNum),
		Tags:	  tags,
	}, nil
}

func (m *Media) ConvertToPersistence(media domain.Media, filename string) *Media{
	if m == nil { m = &Media{} }
	return &Media{
		Id:       uuid.NewV4().String(),
		Type:     media.Type,
		PostId:   media.PostId,
		Filename: filename,
		OrderNum: int(media.OrderNum),
	}
}

func (t Tag) ConvertToDomain(username string) domain.Tag {
	return domain.Tag{
		MediaId:   t.MediaId,
		UserId:    t.UserId,
		Username:  username,
	}
}

func (t *Tag) ConvertToPersistence(tag domain.Tag) *Tag{
	if t == nil { t = &Tag{} }
	return &Tag{
		MediaId:    tag.MediaId,
		UserId:    	tag.UserId,
	}
}

func (f Favorites) ConvertToDomain(collections []domain.Collection, unclassified []domain.ReducedPost) domain.Favorites {
	return domain.Favorites{
		UserId:       f.UserId,
		Collections:  collections,
		Unclassified: unclassified,
	}
}

func (f *Favorites) ConvertToPersistence(favorites domain.FavoritesRequest) *Favorites{
	if f == nil { f = &Favorites{} }
	return &Favorites{
		PostId:       favorites.PostId,
		CollectionId: favorites.CollectionId,
		UserId: 	  favorites.UserId,
	}
}

func (c Collection) ConvertToDomain(posts []domain.ReducedPost) domain.Collection {
	return domain.Collection{
		Id:     c.Id,
		Name:   c.Name,
		UserId: c.UserId,
		Posts:  posts,
	}
}

func (c *Collection) ConvertToPersistence(collection domain.Collection) *Collection{
	if c == nil { c = &Collection{} }
	return &Collection{
		Id:     uuid.NewV4().String(),
		Name:   collection.Name,
		UserId: collection.UserId,
	}
}
