package view

import (
	"attendance-record/client/model"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var window fyne.Window

func NewHistoryWindow(a *model.AppContainer) {
	if window == nil {
		window = a.App.NewWindow("打刻履歴")
	}

	curDtData := binding.NewString()
	vm := viewmodel.NewHistoryViewModel(a, window, curDtData)

	curDtLabel := widget.NewLabelWithData(curDtData)
	curDtLabel.TextStyle = fyne.TextStyle{Bold: true}
	table := NewHistoryListView(vm)
	toolbar := NewHistoryToolbarView(vm)

	window.SetContent(container.New(
		layout.NewBorderLayout(toolbar, curDtLabel, nil, nil),
		toolbar,
		curDtLabel,
		table,
	))

	window.SetCloseIntercept(func() { window.Hide() })
	window.Resize(fyne.NewSize(550, 500))
	window.Show()
}
