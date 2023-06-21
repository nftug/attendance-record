package entity

import (
	"attendance-record/domain/dto"
	"attendance-record/shared/util"
	"errors"
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

func (ts *TimeStatus) Edit(cmd dto.TimeStatusCommandDto) error {
	if cmd.EndedOn == *new(time.Time) &&
		util.GetDate(ts.StartTime) != util.GetDate(time.Now()) {
		return errors.New("cannot set blank endTime")
	}

	if cmd.EndedOn != *new(time.Time) &&
		cmd.StartedOn.Unix() > cmd.EndedOn.Unix() {
		return errors.New("startTime is larger than endTime")
	}

	ts.StartTime = cmd.StartedOn
	ts.EndTime = cmd.EndedOn

	return nil
}

func (ts TimeStatus) TotalTime(now time.Time) time.Duration {
	if ts.IsActive() {
		return now.Sub(ts.StartTime)
	} else {
		return ts.EndTime.Sub(ts.StartTime)
	}
}
