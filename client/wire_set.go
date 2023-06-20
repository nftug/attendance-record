package client

import (
	"attendance-record/client/model"
	"attendance-record/shared"

	"github.com/google/wire"
)

var localApiSet = wire.NewSet(shared.Set, model.NewLocalApi)

var Set = wire.NewSet(
	localApiSet,
	model.NewTimeStatusReceiverSingleton,
	NewClient,
)
