package view

import (
	"attendance-record/client/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewTimeStatusView(w fyne.Window, r *model.TimeStatusReceiver) *fyne.Container {
	return container.NewVBox(
		NewClockView(),
		NewCommandsView(r, w),
		NewStatusView(r),
	)
}
