package view

import (
	"attendance-record/client/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewWindowContent(api *model.Api, w fyne.Window) *fyne.Container {
	return container.NewVBox(NewClock(), NewCommands(api, w), NewStatus(api))
}
