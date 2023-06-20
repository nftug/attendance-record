package model

import (
	"attendance-record/domain/dto"
	"attendance-record/shared"
	"attendance-record/usecase"
)

type timeStatusLocalApi struct {
	usecase *usecase.TimeStatusUseCase
}

func NewLocalApi(a *shared.App) ITimeStatusApi {
	return &timeStatusLocalApi{usecase: a.TimeStatusUseCase}
}

func (api *timeStatusLocalApi) ToggleWork() dto.CurrentTimeStatusDto {
	return api.usecase.ToggleWork()
}

func (api *timeStatusLocalApi) ToggleRest() dto.CurrentTimeStatusDto {
	return api.usecase.ToggleRest()
}

func (api *timeStatusLocalApi) GetCurrentStatus() dto.CurrentTimeStatusDto {
	return api.usecase.GetCurrent()
}
