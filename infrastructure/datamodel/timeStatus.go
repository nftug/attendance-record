package datamodel

import (
	"attendance-record/domain/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TimeStatus struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	StartTime time.Time
	EndTime   time.Time
}

func NewTimeStatusFromEntity(e entity.TimeStatus) TimeStatus {
	return TimeStatus{ID: e.Id, StartTime: e.StartTime, EndTime: e.EndTime}
}

func (d *TimeStatus) ToEntity() entity.TimeStatus {
	return entity.TimeStatus{Id: d.ID, StartTime: d.StartTime, EndTime: d.EndTime}
}

type WorkTimeStatus TimeStatus
type RestTimeStatus TimeStatus
