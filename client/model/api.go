package model

import (
	"attendance-record/domain/dto"
	"attendance-record/shared"
	"attendance-record/usecase"
)

type Api struct {
	usecase *usecase.TimeStatusUseCase
}

func NewApi(a *shared.App) *Api {
	return &Api{usecase: a.TimeStatusUseCase}
}

func (api *Api) ToggleWork() dto.CurrentTimeStatusDto {
	return api.usecase.ToggleWork()
}

func (api *Api) ToggleRest() dto.CurrentTimeStatusDto {
	return api.usecase.ToggleRest()
}

func (api *Api) GetCurrentStatus() dto.CurrentTimeStatusDto {
	return api.usecase.GetCurrent()
}
