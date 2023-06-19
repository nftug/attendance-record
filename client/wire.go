//go:build wireinject
// +build wireinject

package client

import (
	"attendance-record/client/view"
	"attendance-record/shared"

	"fyne.io/fyne/v2"
	"github.com/google/wire"
)

func initTimeStatusView(w fyne.Window, a *shared.App) *view.TimeStatusView {
	wire.Build(Set)
	return nil
}
