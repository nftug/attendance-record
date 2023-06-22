package view

import (
	"attendance-record/client/viewmodel"
	"attendance-record/domain/dto"
	"attendance-record/domain/enum"
	"attendance-record/shared/util"
	"time"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowEditDialog(vm *viewmodel.HistoryViewModel) error {
	item, err := vm.GetSelected()
	if err != nil {
		return err
	}

	lDate := widget.NewLabel(util.GetDate(item.StartedOn).Format(dto.DateFormat))
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

func ShowCreateDialog(vm *viewmodel.HistoryViewModel) (err error) {
	dateEntry := widget.NewEntry()
	startEntry := widget.NewEntry()
	endEntry := widget.NewEntry()

	now := time.Now()
	if vm.CurDt.Year() == now.Year() && vm.CurDt.Month() == now.Month() {
		dateEntry.Text = now.Format(dto.DateFormat)
	} else {
		dateEntry.Text = vm.CurDt.Format(dto.DateFormat)
	}

	startEntry.Text = "00:00"
	endEntry.Text = "00:00"

	var typeSelected enum.TimeStatusType
	typeSelect := widget.NewSelect(
		[]string{"勤務時間", "休憩時間"},
		func(v string) {
			if v == "勤務時間" {
				typeSelected = enum.Work
			} else if v == "休憩時間" {
				typeSelected = enum.Rest
			}
		},
	)
	typeSelect.SetSelected("勤務時間")

	form := widget.NewForm(
		widget.NewFormItem("記録日", dateEntry),
		widget.NewFormItem("種類", typeSelect),
		widget.NewFormItem("開始時刻", startEntry),
		widget.NewFormItem("終了時刻", endEntry),
	)

	dialog.ShowCustomConfirm("記録の新規作成", "保存", "キャンセル", form,
		func(ans bool) {
			if !ans {
				return
			}
			err = vm.Create(typeSelected, dateEntry.Text, startEntry.Text, endEntry.Text)
		}, vm.Window)
	return
}

func ShowDeleteDialog(vm *viewmodel.HistoryViewModel) error {
	item, err := vm.GetSelected()
	if err != nil {
		return err
	}

	var msg string
	if item.Type == enum.Work {
		msg += "警告！これは勤務時間のデータです！\n\n"
	}
	msg += "選択した記録を削除しますか？"

	dialog.ShowConfirm("記録の削除", msg,
		func(ans bool) {
			if !ans {
				return
			}
			err = vm.Delete(item)
		}, vm.Window)
	return err
}
