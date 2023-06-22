package repository

import (
	"attendance-record/domain/entity"
	"attendance-record/domain/interfaces"
	"attendance-record/infrastructure/datamodel"
	"attendance-record/shared/util"
	"errors"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/google/uuid"
)

func NewWorkDummyRepository() interfaces.IWorkRepository {
	return &timeStatusDummyRepository{}
}

func NewRestDummyRepository() interfaces.IRestRepository {
	return &timeStatusDummyRepository{}
}

type timeStatusDummyRepository struct {
	data []datamodel.TimeStatus
}

func (r *timeStatusDummyRepository) Create(item entity.TimeStatus) error {
	d := datamodel.NewTimeStatusFromEntity(item)
	r.data = append(r.data, d)
	return nil
}

func (r *timeStatusDummyRepository) Update(item entity.TimeStatus) error {
	idx := linq.From(r.data).
		IndexOfT(func(x datamodel.TimeStatus) bool { return x.Id == item.Id })
	if idx == -1 {
		return errors.New("the item with specified id cannot be found")
	}
	r.data[idx] = datamodel.NewTimeStatusFromEntity(item)
	return nil
}

func (r *timeStatusDummyRepository) FindByDate(start time.Time, end time.Time) (linq.Query, error) {
	return linq.From(r.data).
		WhereT(func(x datamodel.TimeStatus) bool { return x.StartTime.After(start) && x.StartTime.Before(end) }).
		OrderByT(orderByPredicate).
		SelectT(toEntitySelector), nil
}

func (r *timeStatusDummyRepository) GetLatest() (*entity.TimeStatus, error) {
	if l, ok := linq.From(r.data).
		WhereT(getWhereDayPredicate(util.GetNowDateTime())).
		OrderByT(orderByPredicate).
		Last().(datamodel.TimeStatus); ok {
		p := l.ToEntity()
		return &p, nil
	} else {
		return nil, nil
	}
}

func (r *timeStatusDummyRepository) Delete(id uuid.UUID) error {
	idx := linq.From(r.data).
		IndexOfT(func(x datamodel.TimeStatus) bool { return x.Id == id })
	r.data = append(r.data[:idx], r.data[idx+1:]...)
	return nil
}

func (r *timeStatusDummyRepository) Get(id uuid.UUID) (*entity.TimeStatus, error) {
	v, ok := linq.From(r.data).
		SelectT(toEntitySelector).
		FirstWithT(func(x entity.TimeStatus) bool { return x.Id == id }).(entity.TimeStatus)
	if ok {
		return &v, nil
	} else {
		return nil, errors.New("cannot get item")
	}
}

func orderByPredicate(x datamodel.TimeStatus) int64 {
	return x.StartTime.Unix()
}

func getWhereDayPredicate(dt time.Time) func(datamodel.TimeStatus) bool {
	today, tomorrow := getDayPair(dt)
	return func(x datamodel.TimeStatus) bool {
		return x.StartTime.After(today) && x.StartTime.Before(tomorrow)
	}
}
