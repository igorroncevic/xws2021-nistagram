package persistence

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	uuid "github.com/satori/go.uuid"
	"time"
)

func (p Post) ConvertToDomain(comments []domain.Comment, likes []domain.Like, dislikes []domain.Like, tags []domain.Tag, media []domain.Media) domain.Post{
	return domain.Post{
		Objava:   domain.Objava{
			Id:          p.Id,
			UserId:      p.UserId,
			IsAd:        p.IsAd,
			Type:        p.Type,
			Description: p.Description,
			Location:    p.Location,
			Tags:        tags,
			CreatedAt:   p.CreatedAt,
			Media:       media,
		},
		Comments: comments,
		Likes:    likes,
		Dislikes: dislikes,
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
