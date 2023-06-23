package view

import (
	"attendance-record/client/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func SetSystemTrayMenu(a *model.AppContainer, w fyne.Window) {
	if desk, ok := a.App.(desktop.App); ok {
		qMenu := fyne.NewMenuItem("終了", a.App.Quit)
		qMenu.IsQuit = true

		m := fyne.NewMenu(
			"勤怠記録",
			fyne.NewMenuItem("表示", w.Show),
			fyne.NewMenuItem("打刻履歴", func() { NewHistoryWindow(a) }),
			fyne.NewMenuItemSeparator(),
			qMenu,
		)
		desk.SetSystemTrayMenu(m)
	}
	w.SetCloseIntercept(w.Hide)
}
