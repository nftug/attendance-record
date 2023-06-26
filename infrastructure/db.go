package infrastructure

import (
	"attendance-record/infrastructure/datamodel"
	"attendance-record/infrastructure/localpath"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instance *gorm.DB

func NewDBSingleton(lp *localpath.LocalPathService) *gorm.DB {
	fn := lp.GetJoinedPath("attendance.db")

	if instance == nil {
		db, err := gorm.Open(sqlite.Open(fn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database")
		}

		db.AutoMigrate(
			&datamodel.WorkTimeStatus{},
			&datamodel.RestTimeStatus{},
		)
		instance = db
	}

	return instance
}
