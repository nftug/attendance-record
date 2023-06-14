package usecase

import (
	"attendance-record/domain/dto"
	"attendance-record/shared"
)

func ToggleWork(session *shared.Session) *dto.CurrentTimeStatusDto {
	session.TimeStatusService.ToggleWork()
	return session.TimeStatusService.GetCurrent()
}

func ToggleRest(session *shared.Session) *dto.CurrentTimeStatusDto {
	session.TimeStatusService.ToggleRest()
	return session.TimeStatusService.GetCurrent()
}

func GetCurrent(session *shared.Session) *dto.CurrentTimeStatusDto {
	return session.TimeStatusService.GetCurrent()
}
