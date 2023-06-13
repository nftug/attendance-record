package client

import (
	"client/view"

	"fyne.io/fyne/v2/app"
)

func ShowAndRun() {
	a := app.New()
	w := a.NewWindow("Clock")

	w.SetContent(view.NewWindowContent())
	w.ShowAndRun()
}
