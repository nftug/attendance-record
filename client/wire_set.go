package client

import (
	"attendance-record/client/model"
	"attendance-record/client/view"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	model.NewLocalApi,
	model.NewTimeStatusReceiverSingleton,
	view.NewTimeStatusView,
)
