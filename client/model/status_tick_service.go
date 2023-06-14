package model

import (
	"attendance-record/domain/dto"
	"time"
)

type StatusTickService struct {
	api       *Api
	model     dto.CurrentTimeStatusDto
	WorkTotal time.Duration
	RestTotal time.Duration
	update    func()
}

func NewStatusTickService(api *Api, update func()) *StatusTickService {
	st := api.GetCurrentStatus()
	s := &StatusTickService{api, st, st.Work.TotalTime, st.Rest.TotalTime, update}
	s.startUpdateTick()
	return s
}

func (s *StatusTickService) startUpdateTick() {
	go func() {
		for range time.Tick(time.Second) {
			s.model = s.api.GetCurrentStatus()
			onTickTimer(s.model.Work, &s.WorkTotal)
			onTickTimer(s.model.Rest, &s.RestTotal)
			s.update()
		}
	}()
}

func onTickTimer(ts dto.TimeStatusItemDto, d *time.Duration) {
	if !ts.IsActive || !ts.IsToggleEnabled {
		return
	}
	*d += time.Duration(1) * time.Second
}
