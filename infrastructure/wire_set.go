package infrastructure

import (
	"attendance-record/infrastructure/repository"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	// For dummy
	// repository.NewWorkDummyRepository,
	// repository.NewRestDummyRepository,
	repository.NewWorkRepository,
	repository.NewRestRepository,
	NewDBSingleton,
)
