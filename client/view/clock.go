package view

import (
	"client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func NewClock() *canvas.Text {
	clock := canvas.NewText("", theme.ForegroundColor())
	clock.TextSize = 72
	clock.Alignment = fyne.TextAlignCenter

	viewmodel.UpdateByTick(func(v string) {
		clock.Text = v
		clock.Refresh()
	})

	return clock
}
