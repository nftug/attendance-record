package view

import (
	"attendance-record/client/model"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var preferenceWindow fyne.Window

func ShowPreferenceWindow(app *model.AppContainer) {
	if preferenceWindow == nil {
		preferenceWindow = fyne.CurrentApp().NewWindow("設定")
	}

	vm := viewmodel.NewPreferenceViewModel(app, preferenceWindow)
	tabs := NewPreferenceTabView(vm)
	buttons := NewPreferenceButtonView(vm)

	preferenceWindow.SetContent(container.New(
		layout.NewBorderLayout(nil, buttons, nil, nil),
		tabs,
		buttons,
	))

	preferenceWindow.SetCloseIntercept(preferenceWindow.Hide)
	preferenceWindow.Resize(fyne.NewSize(500, 400))

	preferenceWindow.Show()
}
