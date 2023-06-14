package viewmodel

import (
	"attendance-record/client/model"
	"fmt"
)

type StatusViewModel struct {
	tickService *model.StatusTickService
	WorkTotal   Binding[string]
	RestTotal   Binding[string]
}

func (vm *StatusViewModel) update() {
	vm.WorkTotal.Set(fmt.Sprintf("総勤務時間: %s", vm.tickService.WorkTotal))
	vm.RestTotal.Set(fmt.Sprintf("総休憩時間: %s ", vm.tickService.RestTotal))
}

func NewStatusViewModel(api *model.Api, work Binding[string], rest Binding[string]) *StatusViewModel {
	vm := &StatusViewModel{nil, work, rest}
	vm.tickService = model.NewStatusTickService(api, vm.update)
	vm.update()
	return vm
}
