package viewmodel

import (
	"attendance-record/client/model"
	"attendance-record/shared/appinfo"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CommandsViewModel struct {
	api        model.ITimeStatusApi
	receiver   *model.TimeStatusReceiver
	btnWorking *widget.Button
	btnResting *widget.Button
	window     fyne.Window
}

func NewCommandsViewModel(
	app *model.AppContainer,
	btnW,
	btnR *widget.Button,
	w fyne.Window,
) *CommandsViewModel {
	vm := &CommandsViewModel{app.Api, app.Receiver, btnW, btnR, w}
	vm.receiver.AddUpdateFunc(vm.updateView)
	vm.updateView()

	btnW.OnTapped = vm.OnPressBtnWorking
	btnR.OnTapped = vm.OnPressBtnResting

	return vm
}

func (vm *CommandsViewModel) OnPressBtnWorking() {
	if vm.receiver.Status.Work.IsActive {
		dialog.ShowConfirm("退勤", "退勤しますか？", func(a bool) {
			if a {
				vm.receiver.ToggleWork()
				vm.ShowCurrentStatusDialog()
			}
		}, vm.window)
	} else {
		vm.receiver.ToggleWork()
	}
}

func (vm *CommandsViewModel) OnPressBtnResting() {
	vm.receiver.ToggleRest()
}

func (vm *CommandsViewModel) ShowCurrentStatusDialog() {
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
		vm.btnWorking.SetIcon(theme.LogoutIcon())
	} else {
		vm.btnWorking.SetText("出勤")
		vm.btnWorking.SetIcon(theme.LoginIcon())
	}

	if s.Rest.IsActive {
		vm.btnResting.SetText("休憩終了")
	} else {
		vm.btnResting.SetText("休憩開始")
	}

	if s.Rest.IsActive {
		vm.window.SetTitle(appinfo.AppTitle + " - [休憩中]")
	} else if s.Work.IsActive {
		vm.window.SetTitle(appinfo.AppTitle + " - [勤務中]")
	} else {
		vm.window.SetTitle(appinfo.AppTitle)
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
