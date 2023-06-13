package view

import (
	"client/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewWindowContent() *fyne.Container {
	api := model.NewApi()
	return container.NewVBox(NewClock(), NewCommands(api), NewStatus(api))
}
