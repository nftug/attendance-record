package viewmodel

import (
	"attendance-record/client/model"

	"fyne.io/fyne/v2/data/binding"
)

type StatusViewModel struct {
	receiver  *model.TimeStatusReceiver
	WorkTotal binding.String
	RestTotal binding.String
	OverTime  binding.String
}

func (vm *StatusViewModel) update() {
	workTotal := vm.receiver.Status.Work.TotalTime
	restTotal := vm.receiver.Status.Rest.TotalTime
	overTime := model.Config.OverTime(workTotal)

	vm.WorkTotal.Set("総勤務時間: " + workTotal.String())
	vm.RestTotal.Set("総休憩時間: " + restTotal.String())
	vm.OverTime.Set("残業時間: " + overTime.String())
}

func NewStatusViewModel(app *model.AppContainer) *StatusViewModel {
	vm := &StatusViewModel{
		receiver:  app.Receiver,
		WorkTotal: binding.NewString(),
		RestTotal: binding.NewString(),
		OverTime:  binding.NewString(),
	}
	vm.receiver.AddUpdateFunc(vm.update)
	vm.update()
	return vm
}
