package view

import (
	"attendance-record/client/model"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewStatusView(a *model.AppContainer) fyne.CanvasObject {
	vm := viewmodel.NewStatusViewModel(a)
	form := widget.NewForm(
		widget.NewFormItem("勤務時間", widget.NewLabelWithData(vm.WorkTotal)),
		widget.NewFormItem("休憩時間", widget.NewLabelWithData(vm.RestTotal)),
		widget.NewFormItem("残業時間", widget.NewLabelWithData(vm.Overtime)),
	)
	return container.NewPadded(form)
}
