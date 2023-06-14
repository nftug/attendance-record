package viewmodel

import (
	"client/model"
	"domain/dto"
)

type CommandsViewModel struct {
	api        *model.Api
	model      *dto.TimeStatusSetDto
	btnWorking Button
	btnResting Button
	window     Window
}

func NewCommandsViewModel(api *model.Api, btnW Button, btnR Button, w Window) *CommandsViewModel {
	st := api.LoadTimeStatus()
	vm := &CommandsViewModel{api, st, btnW, btnR, w}
	vm.updateView()
	return vm
}

func (vm *CommandsViewModel) OnPressBtnWorking() {
	vm.model = vm.api.ToggleWork()
	vm.updateView()
}

func (vm *CommandsViewModel) OnPressBtnResting() {
	vm.model = vm.api.ToggleRest()
	vm.updateView()
}

func (vm *CommandsViewModel) updateView() {
	vm.updateByIsActive()
	vm.updateByBtnEnabled()
}

func (vm *CommandsViewModel) updateByIsActive() {
	if vm.model.Work.IsActive {
		vm.btnWorking.SetText("退勤")
	} else {
		vm.btnWorking.SetText("出勤")
		vm.window.SetTitle("勤怠記録")
	}

	if vm.model.Rest.IsActive {
		vm.btnResting.SetText("休憩終了")
	} else {
		vm.btnResting.SetText("休憩開始")
	}

	if vm.model.Rest.IsActive {
		vm.window.SetTitle("勤怠記録 - [休憩中]")
	} else if vm.model.Work.IsActive {
		vm.window.SetTitle("勤怠記録 - [出勤中]")
	} else {
		vm.window.SetTitle("勤怠記録")
	}
}

func (vm *CommandsViewModel) updateByBtnEnabled() {
	if vm.model.Work.IsToggleEnabled {
		vm.btnWorking.Enable()
	} else {
		vm.btnWorking.Disable()
	}

	if vm.model.Rest.IsToggleEnabled {
		vm.btnResting.Enable()
	} else {
		vm.btnResting.Disable()
	}
}
