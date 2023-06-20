package shared

import (
	"attendance-record/domain"
	"attendance-record/infrastructure"
	"attendance-record/usecase"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewAppSingleton,
	domain.Set,
	infrastructure.Set,
	usecase.Set,
)
