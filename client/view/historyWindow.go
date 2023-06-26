package view

import (
	"attendance-record/client/model"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var historyWindow fyne.Window

func ShowHistoryWindow(a *model.AppContainer) {
	if historyWindow == nil {
		historyWindow = a.App.NewWindow("打刻履歴")
	}

	vm := viewmodel.NewHistoryViewModel(a, historyWindow)

	table := NewHistoryListView(vm)
	toolbar := NewHistoryToolbarView(vm)
	footer := NewHistoryFooterView(vm)

	historyWindow.SetContent(container.New(
		layout.NewBorderLayout(toolbar, footer, nil, nil),
		toolbar,
		footer,
		table,
	))

	historyWindow.SetCloseIntercept(historyWindow.Hide)
	historyWindow.Resize(fyne.NewSize(650, 500))
	historyWindow.Show()
}
