package view

import (
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewHistoryToolbarView(vm *viewmodel.HistoryViewModel) fyne.CanvasObject {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() { vm.PrevMonth() }),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() { vm.NextMonth() }),
		widget.NewToolbarAction(theme.HomeIcon(), func() { vm.CurrentMonth() }),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() { vm.InvokeUpdate() }),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() { ShowCreateDialog(vm) }),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() { ShowEditDialog(vm) }),
		widget.NewToolbarAction(theme.ContentClearIcon(), func() { ShowDeleteDialog(vm) }),
	)
}
