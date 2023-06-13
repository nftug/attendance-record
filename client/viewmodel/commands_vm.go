package viewmodel

import (
	"domain/entity"
	"usecase"
)

type CommandsViewModel struct {
	*entity.TimeStatusSet
	btnWorking CommandButton
	btnResting CommandButton
}

type CommandButton interface {
	SetText(v string)
	Enable()
	Disable()
}

func NewCommandsViewModel(
	tss *entity.TimeStatusSet,
	btnWorking CommandButton,
	btnResting CommandButton,
) *CommandsViewModel {
	vm := &CommandsViewModel{
		TimeStatusSet: tss,
		btnWorking:    btnWorking,
		btnResting:    btnResting,
	}
	vm.updateView()
	return vm
}

func (vm *CommandsViewModel) OnPressBtnWorking() {
	usecase.ToggleWork(vm.TimeStatusSet)
	vm.updateView()
}

func (vm *CommandsViewModel) OnPressBtnResting() {
	usecase.ToggleRest(vm.TimeStatusSet)
	vm.updateView()
}

func (vm *CommandsViewModel) updateView() {
	vm.updateBtnText()
	vm.updateBtnEnabled()
}

func (vm *CommandsViewModel) updateBtnText() {
	if vm.Work.IsActive {
		vm.btnWorking.SetText("Leave")
	} else {
		vm.btnWorking.SetText("Attend")
	}

	if vm.Rest.IsActive {
		vm.btnResting.SetText("End Rest")
	} else {
		vm.btnResting.SetText("Start Rest")
	}
}

func (vm *CommandsViewModel) updateBtnEnabled() {
	if vm.Work.IsToggleEnabled {
		vm.btnWorking.Enable()
	} else {
		vm.btnWorking.Disable()
	}

	if vm.Rest.IsToggleEnabled {
		vm.btnResting.Enable()
	} else {
		vm.btnResting.Disable()
	}
}
