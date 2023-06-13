package viewmodel

import (
	"fmt"
	"time"
)

type CommandsViewModel struct {
	isWorking     bool
	isResting     bool
	workStartTime time.Time
	restStartTime time.Time
	workTotalTime time.Duration
	restTotalTime time.Duration

	btnWorking CommandButton
	btnResting CommandButton
}

type CommandButton interface {
	SetText(v string)
	Enable()
	Disable()
}

func NewCommandsViewModel(btnWorking CommandButton, btnResting CommandButton) *CommandsViewModel {
	vm := &CommandsViewModel{btnWorking: btnWorking, btnResting: btnResting}
	vm.updateView()
	return vm
}

func (vm *CommandsViewModel) OnPressBtnWorking() {
	vm.isWorking = !vm.isWorking
	vm.updateView()

	if !vm.isWorking {
		fmt.Printf("%s\n", vm.workTotalTime)
	}
}

func (vm *CommandsViewModel) OnPressBtnResting() {
	vm.isResting = !vm.isResting
	vm.updateView()
}

func (vm *CommandsViewModel) updateView() {
	if vm.isWorking {
		vm.workTotalTime = 0
		vm.workStartTime = time.Now()

		vm.btnWorking.SetText("Leave")
		vm.btnResting.Enable()

	} else {
		vm.workTotalTime += time.Since(vm.workStartTime)
		vm.workStartTime = time.Time{}
		vm.isResting = false

		vm.btnWorking.SetText("Attend")
		vm.btnResting.Disable()
	}

	if vm.isResting {
		vm.restStartTime = time.Now()

		vm.btnResting.SetText("End Rest")
		vm.btnWorking.Disable()
	} else {
		vm.restTotalTime += time.Since(vm.restStartTime)
		vm.restStartTime = time.Time{}

		vm.btnResting.SetText("Start Rest")
		vm.btnWorking.Enable()
	}
}
