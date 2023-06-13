package entity

import (
	"domain/dto"
	"domain/vo"
	"time"
)

type TimeStatusSet struct {
	work vo.TimeStatus
	rest vo.TimeStatus
}

func NewTimeStatusSet() *TimeStatusSet {
	return &TimeStatusSet{
		work: vo.TimeStatus{IsToggleEnabled: true},
		rest: vo.TimeStatus{IsToggleEnabled: false},
	}
}

func (tss *TimeStatusSet) ToggleWork() {
	tss.work.ToggleActive()

	if tss.work.IsActive {
		tss.rest.IsToggleEnabled = true
	} else {
		tss.rest = vo.TimeStatus{IsToggleEnabled: false}
	}
}

func (tss *TimeStatusSet) ToggleRest() {
	tss.rest.ToggleActive()
	tss.work.IsToggleEnabled = !tss.rest.IsActive

	if !tss.rest.IsActive {
		// 休憩が終わったら、総勤務時間から休憩時間を引く
		tss.work.TotalTime -= time.Since(tss.rest.StartTime)
	}
}

func (tss *TimeStatusSet) ToDto() *dto.TimeStatusSetDto {
	return &dto.TimeStatusSetDto{Work: tss.work, Rest: tss.rest}
}
