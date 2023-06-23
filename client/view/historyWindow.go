package view

import (
	"attendance-record/client/model"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var window fyne.Window

func NewHistoryWindow(a *model.AppContainer) {
	if window == nil {
		window = a.App.NewWindow("打刻履歴")
	}

	vm := viewmodel.NewHistoryViewModel(a, window)

	table := NewHistoryListView(vm)
	toolbar := NewHistoryToolbarView(vm)
	footer := NewHistoryFooterView(vm)

	window.SetContent(container.New(
		layout.NewBorderLayout(toolbar, footer, nil, nil),
		toolbar,
		footer,
		table,
	))

	window.SetCloseIntercept(window.Hide)
	window.Resize(fyne.NewSize(650, 500))
	window.Show()
}
