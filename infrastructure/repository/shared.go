package repository

import (
	"attendance-record/domain/entity"
	"attendance-record/infrastructure/datamodel"
	"time"
)

func toEntitySelector(x datamodel.TimeStatus) entity.TimeStatus {
	return x.ToEntity()
}

func getDayPair(dt time.Time) (today time.Time, tomorrow time.Time) {
	today = time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, time.Local)
	tomorrow = today.AddDate(0, 0, 1)
	return
}
