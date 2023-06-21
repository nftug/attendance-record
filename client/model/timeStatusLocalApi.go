package model

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/enum"
	"attendance-record/shared"
	"attendance-record/usecase"

	"github.com/google/uuid"
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

func (api *timeStatusLocalApi) GetAll() []dto.TimeStatusDto {
	return api.usecase.GetAll()
}

func (api *timeStatusLocalApi) Delete(t enum.TimeStatusType, id uuid.UUID) error {
	return api.usecase.Delete(t, id)
}

func (api *timeStatusLocalApi) Update(t enum.TimeStatusType, id uuid.UUID, cmd dto.TimeStatusCommandDto) error {
	return api.usecase.Update(t, id, cmd)
}
