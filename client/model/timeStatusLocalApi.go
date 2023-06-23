package model

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/enum"
	"attendance-record/shared"
	"attendance-record/usecase"
	"time"

	"github.com/google/uuid"
)

type timeStatusLocalApi struct {
	usecase *usecase.TimeStatusUseCase
}

func NewTimeStatusLocalApi(a *shared.App) ITimeStatusApi {
	return &timeStatusLocalApi{usecase: a.TimeStatusUseCase}
}

func (api *timeStatusLocalApi) ToggleWork() error {
	return api.usecase.ToggleWork()
}

func (api *timeStatusLocalApi) ToggleRest() error {
	return api.usecase.ToggleRest()
}

func (api *timeStatusLocalApi) GetCurrentStatus() (*dto.CurrentTimeStatusDto, error) {
	return api.usecase.GetCurrent()
}

func (api *timeStatusLocalApi) FindByMonth(year int, month time.Month) ([]dto.TimeStatusDto, error) {
	return api.usecase.FindByMonth(year, month)
}

func (api *timeStatusLocalApi) Delete(t enum.TimeStatusType, id uuid.UUID) error {
	return api.usecase.Delete(t, id)
}

func (api *timeStatusLocalApi) Update(t enum.TimeStatusType, id uuid.UUID, cmd dto.TimeStatusCommandDto) error {
	return api.usecase.Update(t, id, cmd)
}

func (api *timeStatusLocalApi) Create(t enum.TimeStatusType, cmd dto.TimeStatusCommandDto) error {
	return api.usecase.Create(t, cmd)
}

func (api *timeStatusLocalApi) GetOvertimeByMonth(year int, month time.Month) (*time.Duration, error) {
	return api.usecase.GetOvertimeByMonth(year, month)
}
