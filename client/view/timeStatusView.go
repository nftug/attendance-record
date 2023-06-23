package view

import (
	"attendance-record/client/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewTimeStatusView(a *model.AppContainer, w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		NewClockView(),
		NewCommandsView(a, w),
		NewStatusView(a),
	)
}
