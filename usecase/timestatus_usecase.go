package usecase

import (
	"domain/dto"
	"shared"
)

func ToggleWork(session *shared.Session) *dto.TimeStatusSetDto {
	session.TimeStatusSet.ToggleWork()
	return session.TimeStatusSet.ToDto()
}

func ToggleRest(session *shared.Session) *dto.TimeStatusSetDto {
	session.TimeStatusSet.ToggleRest()
	return session.TimeStatusSet.ToDto()
}

func GetTimeStatus(session *shared.Session) *dto.TimeStatusSetDto {
	// return session.TimeStatusSet.GetCurrent()
	return session.TimeStatusSet.ToDto()
}
