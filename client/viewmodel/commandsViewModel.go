package viewmodel

import (
	"attendance-record/client/model"
	"fmt"
	"time"

	"github.com/multiplay/go-cticker"
)

type CommandsViewModel struct {
	receiver      *model.TimeStatusReceiver
	btnWorking    Button
	btnResting    Button
	btnGetCurrent Button
	window        Window
	fMsg          func(string, string)
}

func NewCommandsViewModel(
	receiver *model.TimeStatusReceiver,
	btnW Button,
	btnR Button,
	btnG Button,
	w Window,
	fMsg func(string, string),
) *CommandsViewModel {
	vm := &CommandsViewModel{receiver, btnW, btnR, btnG, w, fMsg}
	vm.receiver.AddUpdateFunc(vm.updateView)
	vm.updateView()
	return vm
}

func (vm *CommandsViewModel) OnPressBtnWorking() {
	vm.receiver.ToggleWork()
}

func (vm *CommandsViewModel) OnPressBtnResting() {
	vm.receiver.ToggleRest()
}

func (vm *CommandsViewModel) OnPressBtnGetCurrent() {
	vm.receiver.SetCurrentStatus()
	s := vm.receiver.Status
	msg := fmt.Sprintf("勤務時間: %s\n休憩時間: %s\n", s.Work.TotalTime, s.Rest.TotalTime)
	vm.fMsg("取得結果", msg)
}

func (vm *CommandsViewModel) updateView() {
	go vm.updateByIsActive()
	go vm.updateByBtnEnabled()
}

func (vm *CommandsViewModel) updateByIsActive() {
	<-cticker.New(time.Second, 100*time.Millisecond).C
	s := vm.receiver.Status

	if s.Work.IsActive {
		vm.btnWorking.SetText("退勤")
	} else {
		vm.btnWorking.SetText("出勤")
	}

	if s.Rest.IsActive {
		vm.btnResting.SetText("休憩終了")
	} else {
		vm.btnResting.SetText("休憩開始")
	}

	if s.Rest.IsActive {
		vm.window.SetTitle("勤怠記録 - [休憩中]")
	} else if s.Work.IsActive {
		vm.window.SetTitle("勤怠記録 - [出勤中]")
	} else {
		vm.window.SetTitle("勤怠記録")
	}
}

func (vm *CommandsViewModel) updateByBtnEnabled() {
	<-cticker.New(time.Second, 100*time.Millisecond).C
	s := vm.receiver.Status

	if s.Work.IsToggleEnabled {
		vm.btnWorking.Enable()
	} else {
		vm.btnWorking.Disable()
	}

	if s.Rest.IsToggleEnabled {
		vm.btnResting.Enable()
	} else {
		vm.btnResting.Disable()
	}
}
