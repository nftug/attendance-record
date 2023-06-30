package model

import (
	"attendance-record/domain/dto"
	"fmt"
	"strconv"
	"time"

	"github.com/multiplay/go-cticker"
	"github.com/sqweek/dialog"
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

		instance = &TimeStatusReceiver{api, *status, []func(){}, []func(){}}
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
	updateTick := func() {
		for range cticker.New(time.Second, 100*time.Millisecond).C {
			st, err := s.api.GetCurrentStatus()
			if err != nil {
				fmt.Println(err)
			}
			s.Status = *st
			s.invokeUpdate()
		}
	}

	workAlarmTick := func() {
		var workAlarmInvoked, workAlarmSnzInvoked bool
		var workAlarmSnzSec int
		const SnoozeMin = 5

		for range cticker.New(time.Second, 100*time.Millisecond).C {
			if s.Status.Work.EndedOn != *new(time.Time) {
				continue
			}

			if workAlarmSnzInvoked {
				workAlarmSnzSec++
			}

			doWorkAlarm := !workAlarmInvoked && Config.ShouldInvokeWorkAlarm(s.Status.Work.TotalTime)
			doWorkSnooze := workAlarmSnzInvoked && workAlarmSnzSec%(SnoozeMin*60) == 0
			if doWorkAlarm || doWorkSnooze {
				workAlarmInvoked = true

				var msg string
				rem := int(Config.Overtime(s.Status.Work.TotalTime).Minutes() * -1)
				if rem > 0 {
					msg = "あと" + strconv.Itoa(rem) + "分で退勤予定時刻です。"
				} else if rem == 0 {
					msg = "退勤予定時刻になりました。"
				} else {
					msg = "退勤予定時刻を" + strconv.Itoa(rem) + "分超過しています。"
				}

				workAlarmSnzInvoked = dialog.
					Message("%s\n%d分後に再度アラームを表示しますか？", msg, SnoozeMin).Title("勤怠記録").YesNo()
			}
		}
	}

	go updateTick()
	go workAlarmTick()
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
	s.Status = *st
	s.InvokeUpdate()
}
