package entity

import (
	"domain/vo"
)

type TimeStatusSet struct {
	Work *vo.TimeStatus
	Rest *vo.TimeStatus
}

func NewTimeStatusSet() *TimeStatusSet {
	return &TimeStatusSet{
		Work: &vo.TimeStatus{IsToggleEnabled: true},
		Rest: &vo.TimeStatus{IsToggleEnabled: false},
	}
}

func (tss *TimeStatusSet) ToggleWork() {
	tss.Work.ToggleActive()

	if !tss.Work.IsActive {
		tss.Rest = &vo.TimeStatus{IsToggleEnabled: false}
	} else {
		tss.Rest.IsToggleEnabled = true
	}
}

func (tss *TimeStatusSet) ToggleRest() {
	tss.Rest.ToggleActive()
	tss.Work.IsToggleEnabled = !tss.Rest.IsActive

	/*
		if !tss.Rest.IsActive {
			// 休憩が終わったら、総勤務時間から休憩時間を引く
			tss.Work.TotalTime -= time.Since(tss.Rest.StartTime)
		}
	*/
}
