//go:build wireinject
// +build wireinject

package client

import (
	"attendance-record/client/model"
	"attendance-record/client/view"
	"attendance-record/shared"

	"fyne.io/fyne/v2"
	"github.com/google/wire"
)

func initTimeStatusView(w fyne.Window, r *model.TimeStatusReceiver) *view.TimeStatusView {
	wire.Build(view.NewTimeStatusView)
	return nil
}

func initTimeStatusReceiver(a *shared.App) *model.TimeStatusReceiver {
	wire.Build(Set)
	return nil
}
