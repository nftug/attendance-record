package model

import (
	"domain/dto"
	"shared"
	"usecase"
)

type Api struct {
	session *shared.Session
}

func NewApi() *Api {
	return &Api{session: shared.NewSession()}
}

func (api *Api) LoadTimeStatus() *dto.TimeStatusSetDto {
	return usecase.GetTimeStatus(api.session)
}

func (api *Api) ToggleWork() *dto.TimeStatusSetDto {
	return usecase.ToggleWork(api.session)
}

func (api *Api) ToggleRest() *dto.TimeStatusSetDto {
	return usecase.ToggleRest(api.session)
}
