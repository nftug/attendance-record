package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewWindowContent() *fyne.Container {
	return container.NewVBox(NewClock(), NewCommands())
}
