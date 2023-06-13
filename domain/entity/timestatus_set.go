package entity

import "domain/vo"

type TimeStatusSet struct {
	WorkTimeStatus *vo.TimeStatus
	RestTimeStatus *vo.TimeStatus
}

func (tss *TimeStatusSet) ToggleWorkTimeStatus() {
	tss.WorkTimeStatus.ToggleActive()
	tss.RestTimeStatus.IsToggleEnabled = tss.WorkTimeStatus.IsActive
}

func (tss *TimeStatusSet) ToggleRestTimeStatus() {
	tss.RestTimeStatus.ToggleActive()
	tss.WorkTimeStatus.IsToggleEnabled = !tss.RestTimeStatus.IsActive
}
