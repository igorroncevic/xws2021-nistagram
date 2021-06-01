package setup

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"gorm.io/gorm"
)

func FillDatabase(db *gorm.DB) error {
	if db.Migrator().HasTable(&persistence.Post{}){
		db.Migrator().DropTable(&persistence.Post{},
			&persistence.Story{},
			&persistence.Media{},
			&persistence.Tag{},
			&persistence.Collection{},
			&persistence.Favorites{},
			&persistence.PostLikes{},
			&persistence.PostComments{},
			&persistence.HighLights{},
			&persistence.RegistrationRequest{},
			&persistence.Ad{},
			&persistence.Campaign{},
			&persistence.CampaignInfluencerRequest{},
			&persistence.ContentComplaint{},
			&persistence.AdCategory{},
			&persistence.UserAdCategories{},
			&persistence.CampaignChanges{},
		)
	}

	err := db.AutoMigrate(&persistence.Post{},
		&persistence.Story{},
		&persistence.Media{},
		&persistence.Tag{},
		&persistence.Collection{},
		&persistence.Favorites{},
		&persistence.PostLikes{},
		&persistence.PostComments{},
		&persistence.HighLights{},
		&persistence.RegistrationRequest{},
		&persistence.Ad{},
		&persistence.Campaign{},
		&persistence.CampaignInfluencerRequest{},
		&persistence.ContentComplaint{},
		&persistence.AdCategory{},
		&persistence.UserAdCategories{},
		&persistence.CampaignChanges{},
	)

	return err
}
