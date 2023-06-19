package interfaces

import (
	"attendance-record/domain/entity"
	"time"

	"github.com/ahmetb/go-linq/v3"
)

type ITimeStatusRepository interface {
	Create(item entity.TimeStatus)
	Update(item entity.TimeStatus)
	QueryByDate(dt time.Time) linq.Query
	GetLatest() *entity.TimeStatus
}

type IWorkRepository ITimeStatusRepository

type IRestRepository ITimeStatusRepository
