package setup

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model"
	"gorm.io/gorm"
)

func FillDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Post{},
		&model.Story{},
		&model.Media{},
		&model.Tag{},
		&model.Collection{},
		&model.Favorites{},
		&model.PostLikes{},
		&model.PostComments{},
		&model.HighLights{},
		&model.RegistrationRequest{},
		&model.Ad{},
		&model.Campaign{},
		&model.CampaignInfluencerRequest{},
		&model.ContentComplaint{},
		&model.AdCategory{},
		&model.UserAdCategories{},
		&model.CampaignChanges{},
	)

	return err
}
