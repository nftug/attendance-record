package view

import (
	"attendance-record/client/model"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewStatusView(receiver *model.TimeStatusReceiver) *fyne.Container {
	workTotal := binding.NewString()
	restTotal := binding.NewString()
	vm := viewmodel.NewStatusViewModel(receiver, workTotal, restTotal)

	lWorkTotal := widget.NewLabelWithData(any(vm.WorkTotal).(binding.String))
	lRestTotal := widget.NewLabelWithData(any(vm.RestTotal).(binding.String))

	return container.NewVBox(lWorkTotal, lRestTotal)
}
