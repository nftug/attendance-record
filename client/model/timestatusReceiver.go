package model

import (
	"attendance-record/domain/dto"
	"attendance-record/shared/appinfo"
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
		var alarmInvoked, snzInvoked bool
		var snzSec int

		for range cticker.New(time.Second, 100*time.Millisecond).C {
			if s.Status.Work.EndedOn != *new(time.Time) {
				continue
			}
			if snzInvoked {
				snzSec++
			}

			shouldAlarm := !alarmInvoked && Config.ShouldInvokeWorkAlarm(s.Status.Work.TotalTime)
			shouldSnooze := snzInvoked && snzSec%(Config.WorkAlarm.SnoozeMinutes*60) == 0
			if shouldAlarm || shouldSnooze {
				alarmInvoked = true
				var msg string

				rem := int(Config.Overtime(s.Status.Work.TotalTime).Minutes() * -1)
				if rem > 0 {
					msg = "あと" + strconv.Itoa(rem) + "分で退勤予定時刻です。"
				} else if rem == 0 {
					msg = "退勤予定時刻になりました。"
				} else {
					msg = "退勤予定時刻を" + strconv.Itoa(rem) + "分超過しています。"
				}

				snzInvoked = dialog.
					Message("%s\n%d分後に再度アラームを表示しますか？", msg, Config.WorkAlarm.SnoozeMinutes).Title(appinfo.AppTitle).YesNo()
			}
		}
	}

	restAlarmTick := func() {
		var alarmInvoked, snzInvoked bool
		var snzSec int

		for range cticker.New(time.Second, 100*time.Millisecond).C {
			if s.Status.Work.EndedOn != *new(time.Time) {
				continue
			}
			if snzInvoked {
				snzSec++
			}

			shouldAlarm := !alarmInvoked && Config.ShouldInvokeRestAlarm(s.Status.Work.TotalTime, s.Status.Rest.TotalTime)
			shouldSnooze := snzInvoked && snzSec%(Config.RestAlarm.SnoozeMinutes*60) == 0
			if shouldAlarm || shouldSnooze {
				alarmInvoked = true
				snzInvoked = dialog.
					Message("休憩予定時刻になりました。\n%d分後に再度アラームを表示しますか？", Config.RestAlarm.SnoozeMinutes).
					Title(appinfo.AppTitle).YesNo()
			}
		}
	}

	go updateTick()
	go workAlarmTick()
	go restAlarmTick()
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
