package view

import (
	"attendance-record/client/model"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewCommandsView(a *model.AppContainer, w fyne.Window) fyne.CanvasObject {
	btnWorking := widget.NewButton("", func() {})
	btnResting := widget.NewButton("", func() {})
	btnGetCurrent := widget.NewButton("現在の状態", func() {})
	btnHistory := widget.NewButton("履歴", func() { NewHistoryWindow(a) })

	viewmodel.NewCommandsViewModel(a, btnWorking, btnResting, btnGetCurrent, w)

	return container.New(
		layout.NewGridLayoutWithColumns(2),
		btnWorking,
		btnResting,
		btnGetCurrent,
		btnHistory,
	)
}
