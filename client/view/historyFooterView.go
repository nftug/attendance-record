package view

import (
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewHistoryFooterView(vm *viewmodel.HistoryViewModel) fyne.CanvasObject {
	curDtLabel := widget.NewLabelWithData(vm.CurDtData)
	curDtLabel.TextStyle = fyne.TextStyle{Bold: true}
	curDtOvertime := widget.NewLabelWithData(vm.CurDtOvertime)
	return container.NewHBox(curDtLabel, layout.NewSpacer(), curDtOvertime)
}
