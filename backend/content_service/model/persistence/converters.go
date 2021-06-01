package persistence

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
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
