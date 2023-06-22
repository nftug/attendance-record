package startend

import (
	"attendance-record/shared/util"
	"errors"
	"time"
)

type StartEndTime struct {
	StartTime time.Time
	EndTime   time.Time
}

func New(start time.Time, end time.Time) (*StartEndTime, error) {
	if start.After(time.Now()) || end.After(time.Now()) {
		return nil, errors.New("cannot set the future date")
	}
	if end != *new(time.Time) && start.Unix() > end.Unix() {
		return nil, errors.New("start time is larger than end time")
	}
	return &StartEndTime{start, end}, nil
}

func NewAsNow() StartEndTime {
	ret, _ := New(util.GetNowDateTime(), *new(time.Time))
	return *ret
}

func (vo *StartEndTime) Edit(start time.Time, end time.Time) error {
	_, err := New(start, end)
	if err != nil {
		return err
	}
	if end == *new(time.Time) && util.GetDate(vo.StartTime) != util.GetDate(time.Now()) {
		return errors.New("cannot set blank time on end time")
	}

	vo.StartTime = start
	vo.EndTime = end
	return nil
}

func (vo *StartEndTime) IsActive() bool {
	return vo.EndTime == *new(time.Time)
}

func (vo *StartEndTime) TotalTime(now time.Time) time.Duration {
	if vo.IsActive() {
		startDate := util.GetDate(vo.StartTime)
		if startDate == util.GetDate(now) {
			return now.Sub(vo.StartTime)
		} else {
			return startDate.AddDate(0, 0, 1).Sub(vo.StartTime)
		}
	} else {
		return vo.EndTime.Sub(vo.StartTime)
	}
}

func (vo *StartEndTime) SetEnd() {
	vo.EndTime = util.GetNowDateTime()
}
