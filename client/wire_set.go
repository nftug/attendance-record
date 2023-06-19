package client

import (
	"attendance-record/client/model"
	"attendance-record/client/view"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	model.NewApi,
	model.NewTimeStatusReceiverSingleton,
	view.NewTimeStatusView,
	NewClient,
)
