package usecase

import "domain/entity"

func ToggleWorkTimeStatus(tss *entity.TimeStatusSet) {
	tss.ToggleWorkTimeStatus()
}

func ToggleRestTimeStatus(tss *entity.TimeStatusSet) {
	tss.ToggleRestTimeStatus()
}
