package infrastructure

import (
	"attendance-record/infrastructure/datamodel"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instance *gorm.DB

func NewDBSingleton() *gorm.DB {
	if instance == nil {
		db, err := gorm.Open(sqlite.Open("attendance.db"), &gorm.Config{})
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
