package viewmodel

import (
	"client/model"
	"domain/dto"
	"domain/vo"
	"fmt"
	"time"
)

type Binding interface {
	Get() (string, error)
	Set(string) error
}

type StatusViewModel struct {
	api       *model.Api
	model     *dto.TimeStatusSetDto
	workTotal time.Duration
	restTotal time.Duration
	WorkTotal Binding
	RestTotal Binding
}

func (vm *StatusViewModel) update() {
	// w := fmt.Sprintf("Work time total: %s", vm.model.Work.TotalTime)
	// r := fmt.Sprintf("Rest time total: %s ", vm.model.Rest.TotalTime)
	w := fmt.Sprintf("Work time total: %s", vm.workTotal)
	r := fmt.Sprintf("Rest time total: %s ", vm.restTotal)
	vm.WorkTotal.Set(w)
	vm.RestTotal.Set(r)
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

func NewStatusViewModel(api *model.Api, work Binding, rest Binding) *StatusViewModel {
	st := api.LoadTimeStatus()
	vm := &StatusViewModel{api, st, 0, 0, work, rest}
	vm.update()
	vm.startUpdateTick()

	return vm
}

/*
func NewStatusViewModel(api *model.Api, work Binding, rest Binding) *StatusViewModel {
	st := api.LoadTimeStatus()
	vm := &StatusViewModel{api, st, work, rest}
	vm.update()
	vm.startUpdateTick()

	return vm
}
*/
