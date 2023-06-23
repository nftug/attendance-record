package view

import (
	"attendance-record/client/model"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewStatusView(a *model.AppContainer) *fyne.Container {
	vm := viewmodel.NewStatusViewModel(a)
	return container.NewVBox(
		widget.NewLabelWithData(vm.WorkTotal),
		widget.NewLabelWithData(vm.RestTotal),
		widget.NewLabelWithData(vm.OverTime),
	)
}
