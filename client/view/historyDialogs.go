package view

import (
	"attendance-record/client/viewmodel"
	"attendance-record/shared/util"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowEditDialog(vm *viewmodel.HistoryViewModel) {
	if vm.Selected == nil {
		return
	}
	item := *vm.Selected

	startData := binding.NewString()
	endData := binding.NewString()
	startData.Set(util.FormatDateTime(item.StartedOn))
	endData.Set(util.FormatDateTime(item.EndedOn))

	dtLabel := widget.NewLabel(util.GetDate(item.StartedOn).Format("2006-01-02"))
	startEntry := widget.NewEntryWithData(startData)
	endEntry := widget.NewEntryWithData(endData)

	c := container.NewVBox(
		dtLabel,
		widget.NewForm(
			widget.NewFormItem("開始時刻", startEntry),
			widget.NewFormItem("終了時刻", endEntry),
		),
	)

	dialog.ShowCustomConfirm("記録の編集", "保存", "キャンセル", c,
		func(ans bool) {
			if !ans {
				return
			}
			vm.Edit(item, startEntry.Text, endEntry.Text)
		}, vm.Window)
}

func ShowDeleteDialog(vm *viewmodel.HistoryViewModel) {
	if vm.Selected == nil {
		return
	}

	dialog.ShowConfirm("記録の削除", "選択した記録を削除しますか？",
		func(ans bool) {
			if !ans {
				return
			}
			vm.DeleteSelected()
		}, vm.Window)
}
