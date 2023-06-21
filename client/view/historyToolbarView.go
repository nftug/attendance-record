package view

import (
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewHistoryToolbarView(vm *viewmodel.HistoryViewModel) fyne.CanvasObject {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() { ShowEditDialog(vm) }),
		widget.NewToolbarAction(theme.ContentClearIcon(), func() { ShowDeleteDialog(vm) }),
	)
}
