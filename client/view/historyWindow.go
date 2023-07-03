package view

import (
	"attendance-record/client/model"
	"attendance-record/client/resource"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var historyWindow fyne.Window

func ShowHistoryWindow(a *model.AppContainer) {
	if historyWindow == nil {
		historyWindow = fyne.CurrentApp().NewWindow("打刻履歴")
	}

	vm := viewmodel.NewHistoryViewModel(a, historyWindow)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() { vm.PrevMonth() }),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() { vm.NextMonth() }),
		widget.NewToolbarAction(theme.HomeIcon(), func() { vm.CurrentMonth() }),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() { vm.InvokeUpdate() }),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() { ShowCreateDialog(vm) }),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() { ShowEditDialog(vm) }),
		widget.NewToolbarAction(theme.ContentClearIcon(), func() { ShowDeleteDialog(vm) }),
	)
	table := NewHistoryListView(vm)
	footer := NewHistoryFooterView(vm)

	historyWindow.SetContent(container.New(
		layout.NewBorderLayout(toolbar, footer, nil, nil),
		toolbar,
		footer,
		table,
	))

	historyWindow.SetCloseIntercept(historyWindow.Hide)
	historyWindow.Resize(fyne.NewSize(650, 500))
	historyWindow.SetIcon(resource.ResourceIconPng)
	historyWindow.Show()
}

func NewHistoryFooterView(vm *viewmodel.HistoryViewModel) fyne.CanvasObject {
	curDtLabel := widget.NewLabelWithData(vm.CurDtData)
	curDtLabel.TextStyle = fyne.TextStyle{Bold: true}
	curDtOvertime := widget.NewLabelWithData(vm.CurDtOvertime)
	return container.NewHBox(curDtLabel, layout.NewSpacer(), curDtOvertime)
}
