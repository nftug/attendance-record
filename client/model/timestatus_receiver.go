package model

import (
	"attendance-record/domain/dto"
	"time"
)

var instance *TimeStatusReceiver

type TimeStatusReceiver struct {
	api       *Api
	Status    dto.CurrentTimeStatusDto
	WorkTotal time.Duration
	RestTotal time.Duration
	update    []func()
}

func NewTimeStatusReceiverSingleton(api *Api) *TimeStatusReceiver {
	if instance == nil {
		status := api.GetCurrentStatus()
		instance = &TimeStatusReceiver{api, status, status.Work.TotalTime, status.Rest.TotalTime, []func(){}}
		instance.StartUpdateTick()
	}
	return instance
}

func (s *TimeStatusReceiver) AddUpdateFunc(f func()) {
	s.update = append(s.update, f)
}

func (s *TimeStatusReceiver) InvokeUpdate() {
	for _, f := range s.update {
		f()
	}
}

func (s *TimeStatusReceiver) StartUpdateTick() {
	go func() {
		for range time.Tick(time.Second) {
			onTickTimer(s.Status.Work, &s.WorkTotal)
			onTickTimer(s.Status.Rest, &s.RestTotal)
			s.InvokeUpdate()
		}
	}()
}

func onTickTimer(ts dto.TimeStatusItemDto, d *time.Duration) {
	if !ts.IsActive || !ts.IsToggleEnabled {
		return
	}
	*d += time.Duration(1) * time.Second
}

func (s *TimeStatusReceiver) ToggleWork() {
	s.Status = s.api.ToggleWork()
	s.InvokeUpdate()
}

func (s *TimeStatusReceiver) ToggleRest() {
	s.Status = s.api.ToggleRest()
	s.InvokeUpdate()
}

func (s *TimeStatusReceiver) SetCurrentStatus() {
	s.Status = s.api.GetCurrentStatus()
	s.InvokeUpdate()
}
