package view

import (
	"client/viewmodel"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func NewClock() *canvas.Text {
	clock := canvas.NewText("", color.White)
	clock.TextSize = 72
	clock.Alignment = fyne.TextAlignCenter

	viewmodel.UpdateByTick(func(v string) {
		clock.Text = v
		clock.Refresh()
	})

	return clock
}
