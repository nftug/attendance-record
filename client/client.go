package client

import (
	"attendance-record/client/resource"
	"attendance-record/client/view"

	"fyne.io/fyne/v2/app"
)

func Run() {
	a := app.New()
	a.Settings().SetTheme(&resource.MyTheme{})
	w := a.NewWindow("勤怠記録")

	w.SetContent(view.NewWindowContent(w))
	w.ShowAndRun()
}
