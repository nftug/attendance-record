package view

import (
	"attendance-record/client/viewmodel"
	"attendance-record/domain/enum"
	"attendance-record/shared/util"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowEditDialog(vm *viewmodel.HistoryViewModel) error {
	item, err := vm.GetSelected()
	if err != nil {
		return err
	}

	lDate := widget.NewLabel(util.GetDate(item.StartedOn).Format("2006-01-02"))
	startEntry := widget.NewEntry()
	endEntry := widget.NewEntry()
	startEntry.Text = util.FormatDateTime(item.StartedOn)
	endEntry.Text = util.FormatDateTime(item.EndedOn)

	var lType *widget.Label
	if item.Type == enum.Work {
		lType = widget.NewLabel("勤務時間")
	} else {
		lType = widget.NewLabel("休憩時間")
	}

	form := widget.NewForm(
		widget.NewFormItem("記録日", lDate),
		widget.NewFormItem("種類", lType),
		widget.NewFormItem("開始時刻", startEntry),
		widget.NewFormItem("終了時刻", endEntry),
	)

	dialog.ShowCustomConfirm("記録の編集", "保存", "キャンセル", form,
		func(ans bool) {
			if !ans {
				return
			}
			err = vm.Edit(item, startEntry.Text, endEntry.Text)
		}, vm.Window)
	return err
}

func ShowDeleteDialog(vm *viewmodel.HistoryViewModel) error {
	item, err := vm.GetSelected()
	if err != nil {
		return err
	}

	dialog.ShowConfirm("記録の削除", "選択した記録を削除しますか？",
		func(ans bool) {
			if !ans {
				return
			}
			err = vm.Delete(item)
		}, vm.Window)
	return err
}
