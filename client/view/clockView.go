package view

import (
	"attendance-record/client/viewmodel"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func NewClockView() *fyne.Container {
	clock := canvas.NewText("", theme.ForegroundColor())
	clock.TextSize = 72

	viewmodel.UpdateByTick(func(v string) {
		clock.Text = v
		clock.Refresh()
	})

	return container.NewCenter(clock)
}
