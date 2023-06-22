package interfaces

import (
	"attendance-record/domain/entity"
	"attendance-record/domain/enum"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/google/uuid"
)

type ITimeStatusRepository interface {
	Create(item entity.TimeStatus) error
	Update(item entity.TimeStatus) error
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*entity.TimeStatus, error)
	GetLatest() (*entity.TimeStatus, error)
	FindByDate(start time.Time, end time.Time) (linq.Query, error)
}

type IWorkRepository ITimeStatusRepository

type IRestRepository ITimeStatusRepository

type TimeStatusRepositorySet struct {
	workRepository IWorkRepository
	restRepository IRestRepository
}

func NewTimeStatusRepositorySet(wr IWorkRepository, rr IRestRepository) *TimeStatusRepositorySet {
	return &TimeStatusRepositorySet{wr, rr}
}

func (r *TimeStatusRepositorySet) Get(t enum.TimeStatusType) ITimeStatusRepository {
	if t == enum.Work {
		return r.workRepository
	} else if t == enum.Rest {
		return r.restRepository
	}
	return nil
}
