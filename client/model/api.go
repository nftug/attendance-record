package model

import (
	"attendance-record/domain/dto"
	"attendance-record/shared"
	"attendance-record/usecase"
)

type Api struct {
	session *shared.Session
}

func NewApi(session *shared.Session) *Api {
	return &Api{session}
}

func (api *Api) ToggleWork() dto.CurrentTimeStatusDto {
	return usecase.ToggleWork(api.session)
}

func (api *Api) ToggleRest() dto.CurrentTimeStatusDto {
	return usecase.ToggleRest(api.session)
}

func (api *Api) GetCurrentStatus() dto.CurrentTimeStatusDto {
	return usecase.GetCurrent(api.session)
}
