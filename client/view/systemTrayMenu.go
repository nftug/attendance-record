package view

import (
	"attendance-record/client/model"
	"attendance-record/client/resource"
	"attendance-record/shared/appinfo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/sqweek/dialog"
)

func SetSystemTrayMenu(a *model.AppContainer, w fyne.Window) {
	if desk, ok := fyne.CurrentApp().(desktop.App); ok {
		qMenu := fyne.NewMenuItem("終了", fyne.CurrentApp().Quit)
		qMenu.IsQuit = true

		m := fyne.NewMenu(
			appinfo.AppTitle,
			fyne.NewMenuItem("表示", w.Show),
			fyne.NewMenuItem("打刻履歴", func() { ShowHistoryWindow(a) }),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("設定", func() { ShowPreferenceWindow(a) }),
			fyne.NewMenuItem("バージョン情報", func() { showVersionInfo() }),
			fyne.NewMenuItemSeparator(),
			qMenu,
		)
		desk.SetSystemTrayMenu(m)
		desk.SetSystemTrayIcon(resource.ResourceIconPng)
	}
	w.SetCloseIntercept(w.Hide)
}

func showVersionInfo() {
	dialog.
		Message("%s\nver %s\nCreated by %s", appinfo.AppTitle, appinfo.Version, appinfo.Author).
		Title("バージョン情報").Info()
}
