package view

import (
	"attendance-record/shared/appinfo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func ShowVersionInfoDialog(w fyne.Window) {
	w.Show()
	msg := "勤怠記録アプリ (attendance-record)\n" + "ver " + appinfo.Version
	dialog.ShowInformation("バージョン情報", msg, w)
}
