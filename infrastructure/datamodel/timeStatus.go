package datamodel

import (
	"attendance-record/domain/entity"
	"attendance-record/domain/valueobject/startend"
	"time"

	"github.com/google/uuid"
)

type TimeStatus struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;"`
	StartTime time.Time
	EndTime   time.Time
}

func NewTimeStatusFromEntity(e entity.TimeStatus) TimeStatus {
	return TimeStatus{
		Id:        e.Id,
		StartTime: e.Record.StartTime,
		EndTime:   e.Record.EndTime,
	}
}

func (d *TimeStatus) ToEntity() entity.TimeStatus {
	return entity.TimeStatus{
		Id:     d.Id,
		Record: startend.StartEndTime{StartTime: d.StartTime, EndTime: d.EndTime},
	}
}

type WorkTimeStatus TimeStatus
type RestTimeStatus TimeStatus
