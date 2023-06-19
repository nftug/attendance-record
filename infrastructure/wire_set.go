package infrastructure

import (
	"attendance-record/infrastructure/repository"

	"github.com/google/wire"
)

var DummySet = wire.NewSet(
	repository.NewWorkDummyRepository,
	repository.NewRestDummyRepository,
)

var Set = wire.NewSet(
	repository.NewWorkRepository,
	repository.NewRestRepository,
	NewDB,
)
