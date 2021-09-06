package setup

import (
	"github.com/igorroncevic/xws2021-nistagram/monitoring_service/model"
	"gorm.io/gorm"
)

func FillDatabase(db *gorm.DB) error {
	// dropTables(db)

	err := db.AutoMigrate(&model.PerformanceMessage{},
		&model.UserEventMessage{},
	)

	return err
}

func dropTables(db *gorm.DB) {
	if db.Migrator().HasTable(&model.PerformanceMessage{}) {
		db.Migrator().DropTable(&model.PerformanceMessage{},
			&model.UserEventMessage{},
		)
	}
}
