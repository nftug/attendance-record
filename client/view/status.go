package view

import (
	"client/model"
	"client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewStatus(api *model.Api) *fyne.Container {
	workTotal := binding.NewString()
	restTotal := binding.NewString()
	vm := viewmodel.NewStatusViewModel(api, workTotal, restTotal)

	lWorkTotal := widget.NewLabelWithData(vm.WorkTotal.(binding.String))
	lRestTotal := widget.NewLabelWithData(vm.RestTotal.(binding.String))

	return container.NewVBox(lWorkTotal, lRestTotal)
}
