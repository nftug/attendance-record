package view

import (
	"attendance-record/client/model"
	"attendance-record/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewWindowContent(w fyne.Window) *fyne.Container {
	session := shared.NewSession()
	api := model.NewApi(session)
	receiver := model.NewTimeStatusReceiver(api)
	return container.NewVBox(NewClock(), NewCommands(receiver, w), NewStatus(receiver))
}
