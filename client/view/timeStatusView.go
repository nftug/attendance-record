package view

import (
	"attendance-record/client/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type TimeStatusView struct{ *fyne.Container }

func NewTimeStatusView(w fyne.Window, r *model.TimeStatusReceiver) *TimeStatusView {
	c := container.NewVBox(
		NewClockView(),
		NewCommandsView(r, w),
		NewStatusView(r),
	)
	return &TimeStatusView{c}
}
