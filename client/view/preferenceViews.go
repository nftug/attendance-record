package view

import (
	"attendance-record/client/component"
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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
	workAlarmSnoozeEntry := component.NewNumericalEntryWithData(vm.WorkAlarmSnoozeMin)
	workAlarmSetting := widget.NewForm(
		widget.NewFormItem("通知のタイミング",
			container.NewHBox(workAlarmMinEntry, widget.NewLabel("分前")),
		),
		widget.NewFormItem("スヌーズの間隔",
			container.NewHBox(workAlarmSnoozeEntry, widget.NewLabel("分間")),
		),
	)
	vm.WorkAlarmEnabled.AddListener(binding.NewDataListener(func() {
		v, _ := vm.WorkAlarmEnabled.Get()
		if v {
			workAlarmMinEntry.Enable()
			workAlarmSnoozeEntry.Enable()
		} else {
			workAlarmMinEntry.Disable()
			workAlarmSnoozeEntry.Disable()
		}
	}))

	restAlarmCheck := widget.NewCheckWithData("有効にする", vm.RestAlarmEnabled)
	restAlarmHrsEntry := component.NewNumericalEntryWithData(vm.RestAlarmHrs)
	restAlarmMinEntry := component.NewNumericalEntryWithData(vm.RestAlarmMin)
	restAlarmSnoozeEntry := component.NewNumericalEntryWithData(vm.RestAlarmSnoozeMin)
	restAlarmSetting := widget.NewForm(
		widget.NewFormItem("通知のタイミング",
			container.NewHBox(
				widget.NewLabel("勤務開始から"),
				restAlarmHrsEntry, widget.NewLabel("時間"),
				restAlarmMinEntry, widget.NewLabel("分後")),
		),
		widget.NewFormItem("スヌーズの間隔",
			container.NewHBox(restAlarmSnoozeEntry, widget.NewLabel("分間")),
		),
	)
	vm.RestAlarmEnabled.AddListener(binding.NewDataListener(func() {
		v, _ := vm.RestAlarmEnabled.Get()
		if v {
			restAlarmHrsEntry.Enable()
			restAlarmMinEntry.Enable()
			restAlarmSnoozeEntry.Enable()
		} else {
			restAlarmHrsEntry.Disable()
			restAlarmMinEntry.Disable()
			restAlarmSnoozeEntry.Disable()
		}
	}))

	alarmContainer := container.NewVBox(
		widget.NewCard("", "定時前アラーム", container.NewVBox(workAlarmCheck, workAlarmSetting)),
		widget.NewCard("", "休憩前アラーム", container.NewVBox(restAlarmCheck, restAlarmSetting)),
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
