package view

import (
	"attendance-record/client/model"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewCommandsView(a *model.AppContainer, w fyne.Window) fyne.CanvasObject {
	btnWorking := widget.NewButtonWithIcon("", theme.LoginIcon(), func() {})
	btnResting := widget.NewButton("", func() {})
	btnPreference := widget.NewButton("設定", func() { ShowPreferenceWindow(a) })
	btnHistory := widget.NewButton("履歴", func() { ShowHistoryWindow(a) })

	viewmodel.NewCommandsViewModel(a, btnWorking, btnResting, w)

	return container.New(
		layout.NewGridLayoutWithColumns(2),
		btnWorking,
		btnResting,
		btnPreference,
		btnHistory,
	)
}
