package viewmodel

import (
	"attendance-record/client/model"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type CommandsViewModel struct {
	api           model.ITimeStatusApi
	receiver      *model.TimeStatusReceiver
	btnWorking    *widget.Button
	btnResting    *widget.Button
	btnGetCurrent *widget.Button
	window        fyne.Window
}

func NewCommandsViewModel(
	app *model.AppContainer,
	btnW,
	btnR,
	btnGetCurrent *widget.Button,
	w fyne.Window,
) *CommandsViewModel {
	vm := &CommandsViewModel{app.Api, app.Receiver, btnW, btnR, btnGetCurrent, w}
	vm.receiver.AddUpdateFunc(vm.updateView)
	vm.updateView()

	btnW.OnTapped = vm.OnPressBtnWorking
	btnR.OnTapped = vm.OnPressBtnResting
	btnGetCurrent.OnTapped = vm.OnPressBtnGetCurrent

	return vm
}

func (vm *CommandsViewModel) OnPressBtnWorking() {
	if vm.receiver.Status.Work.IsActive {
		dialog.ShowConfirm("退勤", "退勤しますか？", func(a bool) {
			if a {
				vm.receiver.ToggleWork()
				vm.OnPressBtnGetCurrent()
			}
		}, vm.window)
	} else {
		vm.receiver.ToggleWork()
	}
}

func (vm *CommandsViewModel) OnPressBtnResting() {
	vm.receiver.ToggleRest()
}

func (vm *CommandsViewModel) OnPressBtnGetCurrent() {
	vm.receiver.SetCurrentStatus()
	s := vm.receiver.Status
	now := time.Now()

	overtimeTotal, err := vm.api.GetOvertimeByMonth(now.Year(), now.Month())
	if err != nil {
		dialog.ShowError(err, vm.window)
		return
	}

	form := widget.NewForm(
		widget.NewFormItem("勤務時間", widget.NewLabel(s.Work.TotalTime.String())),
		widget.NewFormItem("休憩時間", widget.NewLabel(s.Rest.TotalTime.String())),
		widget.NewFormItem("今日の残業時間", widget.NewLabel(model.Config.Overtime(s.Work.TotalTime).String())),
		widget.NewFormItem("今月の残業時間", widget.NewLabel(overtimeTotal.String())),
	)
	dialog.ShowCustom("本日の勤務記録", "OK", form, vm.window)
}

func (vm *CommandsViewModel) updateView() {
	go vm.updateByIsActive()
	go vm.updateByBtnEnabled()
}

func (vm *CommandsViewModel) updateByIsActive() {
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
