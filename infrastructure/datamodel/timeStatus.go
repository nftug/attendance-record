package datamodel

import (
	"attendance-record/domain/entity"
	"time"

	"github.com/google/uuid"
)

type TimeStatus struct {
	Id        uuid.UUID
	StartTime time.Time
	EndTime   time.Time
}

func NewTimeStatusFromEntity(e entity.TimeStatus) TimeStatus {
	return TimeStatus{e.Id, e.StartTime, e.EndTime}
}

func (d *TimeStatus) ToEntity() entity.TimeStatus {
	return entity.TimeStatus{Id: d.Id, StartTime: d.StartTime, EndTime: d.EndTime}
}
