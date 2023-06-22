package entity

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/valueobject/startend"
	"time"

	"github.com/google/uuid"
)

type TimeStatus struct {
	Id     uuid.UUID
	Record startend.StartEndTime
}

func NewTimeStatusAsNow() TimeStatus {
	return TimeStatus{Id: uuid.New(), Record: startend.NewAsNow()}
}

func NewTimeStatus(cmd dto.TimeStatusCommandDto) (*TimeStatus, error) {
	record, err := startend.New(cmd.StartedOn, cmd.EndedOn)
	if err != nil {
		return nil, err
	}
	return &TimeStatus{Id: uuid.New(), Record: *record}, nil
}

func (ts TimeStatus) IsActive() bool {
	return ts.Record.IsActive()
}

func (ts *TimeStatus) End() {
	ts.Record.SetEnd()
}

func (ts *TimeStatus) Edit(cmd dto.TimeStatusCommandDto) error {
	if err := ts.Record.Edit(cmd.StartedOn, cmd.EndedOn); err != nil {
		return err
	}
	return nil
}

func (ts TimeStatus) TotalTime(now time.Time) time.Duration {
	return ts.Record.TotalTime(now)
}
