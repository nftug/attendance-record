package view

import (
	"attendance-record/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func ShowVersionInfoDialog(w fyne.Window) {
	w.Show()
	msg := "勤怠記録アプリ (attendance-record)\n" + "ver " + shared.Version
	dialog.ShowInformation("バージョン情報", msg, w)
}
