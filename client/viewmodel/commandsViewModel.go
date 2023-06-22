package viewmodel

import (
	"attendance-record/client/model"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/multiplay/go-cticker"
)

type CommandsViewModel struct {
	receiver      *model.TimeStatusReceiver
	btnWorking    *widget.Button
	btnResting    *widget.Button
	btnGetCurrent *widget.Button
	window        fyne.Window
}

func NewCommandsViewModel(
	receiver *model.TimeStatusReceiver,
	btnW *widget.Button,
	btnR *widget.Button,
	btnS *widget.Button,
	w fyne.Window,
) *CommandsViewModel {
	vm := &CommandsViewModel{receiver, btnW, btnR, btnS, w}
	vm.receiver.AddUpdateFunc(vm.updateView)
	vm.updateView()

	btnW.OnTapped = vm.OnPressBtnWorking
	btnR.OnTapped = vm.OnPressBtnResting
	btnS.OnTapped = vm.OnPressBtnSync

	return vm
}

func (vm *CommandsViewModel) OnPressBtnWorking() {
	if vm.receiver.Status.Work.IsActive {
		dialog.ShowConfirm("退勤", "退勤しますか？", func(a bool) {
			if a {
				vm.receiver.ToggleWork()
			}
		}, vm.window)
	} else {
		vm.receiver.ToggleWork()
	}
}

func (vm *CommandsViewModel) OnPressBtnResting() {
	vm.receiver.ToggleRest()
}

func (vm *CommandsViewModel) OnPressBtnSync() {
	vm.receiver.SetCurrentStatus()
	s := vm.receiver.Status
	msg := fmt.Sprintf("勤務時間: %s\n休憩時間: %s\n", s.Work.TotalTime, s.Rest.TotalTime)
	dialog.ShowInformation("同期しました", msg, vm.window)
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
		vm.window.SetTitle("勤怠記録 - [勤務中]")
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
