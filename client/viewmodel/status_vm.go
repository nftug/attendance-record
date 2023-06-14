package viewmodel

import (
	"client/model"
	"domain/dto"
	"domain/vo"
	"fmt"
	"time"
)

type StatusViewModel struct {
	api       *model.Api
	model     *dto.TimeStatusSetDto
	workTotal time.Duration
	restTotal time.Duration
	WorkTotal Binding[string]
	RestTotal Binding[string]
}

func (vm *StatusViewModel) update() {
	vm.WorkTotal.Set(fmt.Sprintf("総勤務時間: %s", vm.workTotal))
	vm.RestTotal.Set(fmt.Sprintf("総休憩時間: %s ", vm.restTotal))
}

func (vm *StatusViewModel) startUpdateTick() {
	go func() {
		for range time.Tick(time.Second) {
			vm.model = vm.api.LoadTimeStatus()
			vm.onTickTimer(vm.model.Work, &vm.workTotal)
			vm.onTickTimer(vm.model.Rest, &vm.restTotal)
			vm.update()
		}
	}()
}

func (vm *StatusViewModel) onTickTimer(ts vo.TimeStatus, d *time.Duration) {
	if !ts.IsActive || !ts.IsToggleEnabled {
		return
	}
	*d += time.Duration(1) * time.Second
}

func NewStatusViewModel(api *model.Api, work Binding[string], rest Binding[string]) *StatusViewModel {
	st := api.LoadTimeStatus()
	vm := &StatusViewModel{api, st, 0, 0, work, rest}
	vm.update()
	vm.startUpdateTick()

	return vm
}
