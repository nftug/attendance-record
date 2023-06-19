package client

import (
	"attendance-record/client/model"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	model.NewLocalApi,
	model.NewTimeStatusReceiverSingleton,
	NewClient,
)
