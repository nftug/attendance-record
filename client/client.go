package client

import (
	"attendance-record/client/model"
	"attendance-record/client/resource"
	"attendance-record/client/view"

	"fyne.io/fyne/v2/app"
)

func ShowAndRun(api *model.Api) {
	a := app.New()
	a.Settings().SetTheme(&resource.MyTheme{})
	w := a.NewWindow("勤怠記録")

	w.SetContent(view.NewWindowContent(api, w))
	w.ShowAndRun()
}
