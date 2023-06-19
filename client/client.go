package client

import (
	"attendance-record/client/resource"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

func Run() {
	a := app.New()
	a.Settings().SetTheme(&resource.MyTheme{})
	w := a.NewWindow("勤怠記録")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu(
			"勤怠記録",
			fyne.NewMenuItem("表示", func() { w.Show() }),
		)
		desk.SetSystemTrayMenu(m)
	}
	w.SetCloseIntercept(func() { w.Hide() })

	v := initTimeStatusView(w)
	w.SetContent(v.Container)
	w.ShowAndRun()
}
