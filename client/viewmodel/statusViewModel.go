package viewmodel

import (
	"attendance-record/client/model"

	"fyne.io/fyne/v2/data/binding"
)

type StatusViewModel struct {
	receiver  *model.TimeStatusReceiver
	WorkTotal binding.String
	RestTotal binding.String
	Overtime  binding.String
}

func NewStatusViewModel(app *model.AppContainer) *StatusViewModel {
	vm := &StatusViewModel{
		receiver:  app.Receiver,
		WorkTotal: binding.NewString(),
		RestTotal: binding.NewString(),
		Overtime:  binding.NewString(),
	}
	vm.receiver.AddUpdateFunc(vm.update)
	vm.update()
	return vm
}

func (vm *StatusViewModel) update() {
	workTotal := vm.receiver.Status.Work.TotalTime
	restTotal := vm.receiver.Status.Rest.TotalTime
	overtime := model.Config.Overtime(workTotal)

	vm.WorkTotal.Set(workTotal.String())
	vm.RestTotal.Set(restTotal.String())
	vm.Overtime.Set(overtime.String())
}
