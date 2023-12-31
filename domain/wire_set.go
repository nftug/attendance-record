package domain

import (
	"attendance-record/domain/interfaces"
	"attendance-record/domain/service"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	service.NewTimeStatusService,
	interfaces.NewTimeStatusRepositorySet,
)
