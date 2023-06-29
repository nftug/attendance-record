package view

import (
	"attendance-record/client/component"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewPreferenceTabView(vm *viewmodel.PreferenceViewModel) fyne.CanvasObject {
	// Work hours
	workHrsSlider := widget.NewSlider(1.0, 8.0)
	workHrsSlider.Value = vm.GetWorkHour()
	workHrsSlider.OnChanged = vm.OnChangeWorkHrsData
	workHrsLabel := widget.NewLabelWithData(vm.WorkHrsLabelData)
	workHrsContainer := container.NewVBox(
		widget.NewForm(widget.NewFormItem("標準の勤務時間", workHrsLabel)),
		workHrsSlider,
	)

	// Alarm
	workAlarmCheck := widget.NewCheckWithData("有効にする", vm.WorkAlarmEnabled)
	workAlarmMinEntry := component.NewNumericalEntryWithData(vm.WorkAlarmBeforeMin)
	workAlarmMin := widget.NewForm(
		widget.NewFormItem("通知のタイミング",
			container.NewHBox(workAlarmMinEntry, widget.NewLabel("分前")),
		))
	alarmContainer := container.NewVBox(
		widget.NewCard("", "退勤前アラーム", container.NewVBox(workAlarmCheck, workAlarmMin)),
	)

	// Local path
	localPathTitleLabel := widget.NewLabel("データの保存先")
	localPathTitleLabel.TextStyle = fyne.TextStyle{Bold: true}
	localPathContainer := container.NewVBox(
		localPathTitleLabel,
		widget.NewLabelWithData(vm.LocalPathData),
		widget.NewButtonWithIcon("開く", theme.FolderOpenIcon(), vm.OpenLocalPath),
	)

	return container.NewAppTabs(
		container.NewTabItem("勤務時間", workHrsContainer),
		container.NewTabItem("アラーム", alarmContainer),
		container.NewTabItem("保存先", localPathContainer),
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
