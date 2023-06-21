package model

import (
	"attendance-record/domain/dto"
	"time"

	"github.com/multiplay/go-cticker"
)

var instance *TimeStatusReceiver

type TimeStatusReceiver struct {
	api         ITimeStatusApi
	Status      dto.CurrentTimeStatusDto
	update      []func()
	updateOuter []func()
}

func NewTimeStatusReceiverSingleton(api ITimeStatusApi) *TimeStatusReceiver {
	if instance == nil {
		status := api.GetCurrentStatus()
		instance = &TimeStatusReceiver{api, status, []func(){}, []func(){}}
		instance.StartUpdateTick()
	}
	return instance
}

func (s *TimeStatusReceiver) AddUpdateFunc(f ...func()) {
	s.update = append(s.update, f...)
}

func (s *TimeStatusReceiver) AddUpdateOuterFunc(f ...func()) {
	s.updateOuter = append(s.updateOuter, f...)
}

func (s *TimeStatusReceiver) invokeUpdate() {
	for _, f := range s.update {
		f()
	}
}

func (s *TimeStatusReceiver) InvokeUpdate() {
	s.invokeUpdate()
	for _, f := range s.updateOuter {
		f()
	}
}

func (s *TimeStatusReceiver) StartUpdateTick() {
	go func() {
		for range cticker.New(time.Second, 100*time.Millisecond).C {
			s.Status = s.api.GetCurrentStatus()
			s.invokeUpdate()
		}
	}()
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
