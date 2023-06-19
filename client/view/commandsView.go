package view

import (
	"attendance-record/client/model"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewCommandsView(receiver *model.TimeStatusReceiver, w fyne.Window) *fyne.Container {
	btnWorking := widget.NewButton("", func() {})
	btnResting := widget.NewButton("", func() {})
	btnGetCurrent := widget.NewButton("集計", func() {})
	btnReset := widget.NewButton("リセット", func() {})
	btnReset.Disable()

	fMsg := func(title string, message string) {
		dialog.ShowInformation(title, message, w)
	}

	vm := viewmodel.NewCommandsViewModel(receiver, btnWorking, btnResting, btnGetCurrent, w, fMsg)
	btnWorking.OnTapped = vm.OnPressBtnWorking
	btnResting.OnTapped = vm.OnPressBtnResting
	btnGetCurrent.OnTapped = vm.OnPressBtnGetCurrent

	return container.New(
		layout.NewGridLayoutWithColumns(2),
		btnWorking,
		btnResting,
		btnGetCurrent,
		btnReset,
	)
}
