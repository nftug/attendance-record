package view

import (
	"client/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewWindowContent(api *model.Api) *fyne.Container {
	return container.NewVBox(NewClock(), NewCommands(api), NewStatus(api))
}
