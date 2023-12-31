package infrastructure

import (
	"attendance-record/infrastructure/localpath"
	"attendance-record/infrastructure/repository"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	// For dummy
	// repository.NewWorkDummyRepository,
	// repository.NewRestDummyRepository,
	repository.NewWorkSqlRepository,
	repository.NewRestSqlRepository,
	NewDBSingleton,
	repository.NewConfigRepository,
	localpath.NewLocalPathService,
)
