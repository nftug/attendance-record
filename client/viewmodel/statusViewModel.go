package viewmodel

import (
	"attendance-record/client/model"
	"fmt"
)

type StatusViewModel struct {
	receiver  *model.TimeStatusReceiver
	WorkTotal Binding[string]
	RestTotal Binding[string]
}

func (vm *StatusViewModel) update() {
	vm.WorkTotal.Set(fmt.Sprintf("総勤務時間: %s", vm.receiver.Status.Work.TotalTime))
	vm.RestTotal.Set(fmt.Sprintf("総休憩時間: %s ", vm.receiver.Status.Rest.TotalTime))
}

func NewStatusViewModel(receiver *model.TimeStatusReceiver, work Binding[string], rest Binding[string]) *StatusViewModel {
	vm := &StatusViewModel{receiver, work, rest}
	vm.receiver.AddUpdateFunc(vm.update)
	vm.update()
	return vm
}
