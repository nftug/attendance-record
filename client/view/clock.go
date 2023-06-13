package view

import (
	"client/viewmodel"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func UpdateClock(clock *canvas.Text) {
	clock.Text = time.Now().Format("03:04:05")
	clock.Refresh()
}

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
