package view

import (
	"attendance-record/client/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func SetSystemTrayMenu(a *model.AppContainer, w fyne.Window) {
	if desk, ok := fyne.CurrentApp().(desktop.App); ok {
		qMenu := fyne.NewMenuItem("終了", fyne.CurrentApp().Quit)
		qMenu.IsQuit = true

		m := fyne.NewMenu(
			"勤怠記録",
			fyne.NewMenuItem("表示", w.Show),
			fyne.NewMenuItem("打刻履歴", func() { ShowHistoryWindow(a) }),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("設定", func() { ShowPreferenceWindow(a) }),
			fyne.NewMenuItem("バージョン情報", func() { ShowVersionInfoDialog(w) }),
			fyne.NewMenuItemSeparator(),
			qMenu,
		)
		desk.SetSystemTrayMenu(m)
	}
	w.SetCloseIntercept(w.Hide)
}
