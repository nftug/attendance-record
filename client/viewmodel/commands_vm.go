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
}

func NewCommandsViewModel(api *model.Api, btnW Button, btnR Button) *CommandsViewModel {
	st := api.LoadTimeStatus()
	vm := &CommandsViewModel{api, st, btnW, btnR}
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
	vm.updateBtnText()
	vm.updateBtnEnabled()
}

func (vm *CommandsViewModel) updateBtnText() {
	if vm.model.Work.IsActive {
		vm.btnWorking.SetText("Leave")
	} else {
		vm.btnWorking.SetText("Attend")
	}

	if vm.model.Rest.IsActive {
		vm.btnResting.SetText("End Rest")
	} else {
		vm.btnResting.SetText("Start Rest")
	}
}

func (vm *CommandsViewModel) updateBtnEnabled() {
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
