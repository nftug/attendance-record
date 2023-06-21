package interfaces

import (
	"attendance-record/domain/entity"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/google/uuid"
)

type ITimeStatusRepository interface {
	Create(item entity.TimeStatus)
	Update(item entity.TimeStatus)
	QueryByDate(dt time.Time) linq.Query
	GetLatest() *entity.TimeStatus
	GetAll() linq.Query
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*entity.TimeStatus, error)
}

type IWorkRepository ITimeStatusRepository

type IRestRepository ITimeStatusRepository
