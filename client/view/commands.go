package view

import (
	"client/model"
	"client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewCommands(api *model.Api, w fyne.Window) *fyne.Container {
	btnWorking := widget.NewButton("", func() {})
	btnResting := widget.NewButton("", func() {})

	vm := viewmodel.NewCommandsViewModel(api, btnWorking, btnResting, w)
	btnWorking.OnTapped = vm.OnPressBtnWorking
	btnResting.OnTapped = vm.OnPressBtnResting

	return container.New(
		layout.NewGridLayoutWithColumns(2),
		btnWorking,
		btnResting,
	)
}
