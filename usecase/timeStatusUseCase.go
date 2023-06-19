package usecase

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/enum"
	"attendance-record/domain/service"
)

type TimeStatusUseCase struct {
	service *service.TimeStatusService
}

func NewTimeStatusUseCase(service *service.TimeStatusService) *TimeStatusUseCase {
	return &TimeStatusUseCase{service: service}
}

func (u *TimeStatusUseCase) ToggleWork() dto.CurrentTimeStatusDto {
	u.service.ToggleState(enum.Work)
	return u.service.GetCurrent()
}

func (u *TimeStatusUseCase) ToggleRest() dto.CurrentTimeStatusDto {
	u.service.ToggleState(enum.Rest)
	return u.service.GetCurrent()
}

func (u *TimeStatusUseCase) GetCurrent() dto.CurrentTimeStatusDto {
	return u.service.GetCurrent()
}
