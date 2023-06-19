package entity

import (
	"attendance-record/shared/util"
	"time"

	"github.com/google/uuid"
)

type TimeStatus struct {
	Id        uuid.UUID
	StartTime time.Time
	EndTime   time.Time // 継続中のフラグを兼ねている
}

func NewTimeStatus() TimeStatus {
	return TimeStatus{Id: uuid.New(), StartTime: util.GetNowDateTime()}
}

func (ts TimeStatus) IsActive() bool {
	return ts.EndTime == *new(time.Time)
}

func (ts *TimeStatus) End() {
	ts.EndTime = util.GetNowDateTime()
}

func (ts TimeStatus) TotalTime(now time.Time) time.Duration {
	if ts.IsActive() {
		return now.Sub(ts.StartTime)
	} else {
		return ts.EndTime.Sub(ts.StartTime)
	}
}
