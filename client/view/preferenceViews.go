package view

import (
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewPreferenceTabView(vm *viewmodel.PreferenceViewModel) fyne.CanvasObject {
	workHrsSlider := widget.NewSlider(1.0, 8.0)
	workHrsSlider.Value = vm.GetWorkHour()
	workHrsSlider.OnChanged = vm.OnChangeWorkHrsData
	workHrsLabel := widget.NewLabelWithData(vm.WorkHrsLabelData)

	workHrsContainer := container.NewVBox(
		widget.NewForm(widget.NewFormItem("標準の勤務時間", workHrsLabel)),
		workHrsSlider,
	)

	return container.NewAppTabs(
		container.NewTabItem("勤務時間", workHrsContainer),
	)
}

func NewPreferenceButtonView(vm *viewmodel.PreferenceViewModel) fyne.CanvasObject {
	return container.NewHBox(
		layout.NewSpacer(),
		widget.NewButtonWithIcon("保存", theme.DocumentSaveIcon(), vm.OnClickSave),
		widget.NewButtonWithIcon("キャンセル", theme.CancelIcon(), vm.OnClickCancel),
		widget.NewButtonWithIcon("適用", theme.ConfirmIcon(), vm.OnClickApply),
	)
}
