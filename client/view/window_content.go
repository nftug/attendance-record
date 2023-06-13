package view

import (
	"usecase"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewWindowContent() *fyne.Container {
	tss := usecase.InitTimeStatusSet()
	return container.NewVBox(NewClock(), NewCommands(tss), NewStatus(tss))
}
