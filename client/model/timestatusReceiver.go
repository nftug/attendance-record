package model

import (
	"attendance-record/domain/dto"
	"fmt"
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
		status, err := api.GetCurrentStatus()
		if err != nil {
			fmt.Println(err)
		}

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
			st, err := s.api.GetCurrentStatus()
			if err != nil {
				fmt.Println(err)
			}
			s.Status = st
			s.invokeUpdate()
		}
	}()
}

func (s *TimeStatusReceiver) ToggleWork() {
	err := s.api.ToggleWork()
	if err != nil {
		fmt.Println(err)
	}
	s.SetCurrentStatus()
}

func (s *TimeStatusReceiver) ToggleRest() {
	err := s.api.ToggleRest()
	if err != nil {
		fmt.Println(err)
	}
	s.SetCurrentStatus()
}

func (s *TimeStatusReceiver) SetCurrentStatus() {
	st, err := s.api.GetCurrentStatus()
	if err != nil {
		fmt.Println(err)
	}
	s.Status = st
	s.InvokeUpdate()
}
