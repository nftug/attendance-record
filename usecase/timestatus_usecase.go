package usecase

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/enum"
	"attendance-record/shared"
)

func ToggleWork(session *shared.Session) dto.CurrentTimeStatusDto {
	session.TimeStatusService.ToggleState(enum.Work)
	return session.TimeStatusService.GetCurrent()
}

func ToggleRest(session *shared.Session) dto.CurrentTimeStatusDto {
	session.TimeStatusService.ToggleState(enum.Rest)
	return session.TimeStatusService.GetCurrent()
}

func GetCurrent(session *shared.Session) dto.CurrentTimeStatusDto {
	return session.TimeStatusService.GetCurrent()
}
