package infrastructure

import (
	"attendance-record/infrastructure/repository"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	repository.NewWorkDummyRepository,
	repository.NewRestDummyRepository,
)
