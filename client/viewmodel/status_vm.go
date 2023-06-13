package viewmodel

import (
	"domain/entity"
	"fmt"
	"time"
)

type Binding interface {
	Get() (string, error)
	Set(string) error
}

type StatusViewModel struct {
	*entity.TimeStatusSet
	WorkTotal Binding
	RestTotal Binding
}

func (vm *StatusViewModel) startUpdate() {
	go func() {
		for range time.Tick(time.Second) {
			if vm.Work.IsActive {
				vm.Work.OnTickTimer()
			}
			vm.Rest.OnTickTimer()

			ws := fmt.Sprintf("%s ", time.Since(vm.Work.StartTime))
			rs := fmt.Sprintf("%s ", time.Since(vm.Rest.StartTime))
			vm.WorkTotal.Set(ws)
			vm.RestTotal.Set(rs)
		}
	}()
}

func NewStatusViewModel(tss *entity.TimeStatusSet, work Binding, rest Binding) *StatusViewModel {
	vm := &StatusViewModel{tss, work, rest}
	vm.startUpdate()

	return vm
}
