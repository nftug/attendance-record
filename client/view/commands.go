package view

import (
	"client/viewmodel"
	"domain/entity"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewCommands(tss *entity.TimeStatusSet) *fyne.Container {
	btnWorking := widget.NewButton("", func() {})
	btnResting := widget.NewButton("", func() {})

	vm := viewmodel.NewCommandsViewModel(tss, btnWorking, btnResting)
	btnWorking.OnTapped = vm.OnPressBtnWorking
	btnResting.OnTapped = vm.OnPressBtnResting

	return container.New(
		layout.NewGridLayoutWithColumns(2),
		btnWorking,
		btnResting,
	)
}
