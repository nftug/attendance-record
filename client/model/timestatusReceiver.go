package model

import (
	"attendance-record/domain/dto"
	"time"

	"github.com/multiplay/go-cticker"
)

var instance *TimeStatusReceiver

type TimeStatusReceiver struct {
	api       ITimeStatusApi
	Status    dto.CurrentTimeStatusDto
	WorkTotal time.Duration
	RestTotal time.Duration
	update    []func()
}

func NewTimeStatusReceiverSingleton(api ITimeStatusApi) *TimeStatusReceiver {
	if instance == nil {
		status := api.GetCurrentStatus()
		instance = &TimeStatusReceiver{api, status, status.Work.TotalTime, status.Rest.TotalTime, []func(){}}
		instance.StartUpdateTick()
	}
	return instance
}

func (s *TimeStatusReceiver) AddUpdateFunc(f ...func()) {
	s.update = append(s.update, f...)
}

func (s *TimeStatusReceiver) InvokeUpdate() {
	for _, f := range s.update {
		f()
	}
}

func (s *TimeStatusReceiver) StartUpdateTick() {
	go func() {
		for range cticker.New(time.Second, 100*time.Millisecond).C {
			onTickTimer(s.Status.Work, &s.WorkTotal)
			onTickTimer(s.Status.Rest, &s.RestTotal)
			s.InvokeUpdate()
		}
	}()
}

func onTickTimer(ts dto.CurrentTimeStatusItemDto, d *time.Duration) {
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
	s.WorkTotal = s.Status.Work.TotalTime
	s.RestTotal = s.Status.Rest.TotalTime
	s.InvokeUpdate()
}
