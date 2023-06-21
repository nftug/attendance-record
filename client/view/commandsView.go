package view

import (
	"attendance-record/client/model"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewCommandsView(a *model.AppContainer, w fyne.Window) *fyne.Container {
	btnWorking := widget.NewButton("", func() {})
	btnResting := widget.NewButton("", func() {})
	btnSync := widget.NewButton("同期", func() {})
	btnHistory := widget.NewButton("履歴", func() { NewHistoryWindow(a) })

	viewmodel.NewCommandsViewModel(a.Receiver, btnWorking, btnResting, btnSync, w)

	return container.New(
		layout.NewGridLayoutWithColumns(2),
		btnWorking,
		btnResting,
		btnSync,
		btnHistory,
	)
}
