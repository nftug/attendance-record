package repository

import (
	"attendance-record/domain/entity"
	"attendance-record/domain/interfaces"
	"attendance-record/infrastructure/datamodel"
	"attendance-record/shared/util"
	"log"
	"time"

	"github.com/ahmetb/go-linq/v3"
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

func (r *timeStatusDummyRepository) Create(item entity.TimeStatus) {
	d := datamodel.NewTimeStatusFromEntity(item)
	r.data = append(r.data, d)
}

func (r *timeStatusDummyRepository) Update(item entity.TimeStatus) {
	idx := linq.From(r.data).IndexOfT(func(x datamodel.TimeStatus) bool { return x.Id == item.Id })
	if idx == -1 {
		log.Fatal("The item with specified id cannot be found.")
	}
	r.data[idx] = datamodel.NewTimeStatusFromEntity(item)
}

func (r *timeStatusDummyRepository) QueryByDate(dt time.Time) linq.Query {
	predicate := getWhereDayPredicate(dt)
	return linq.From(r.data).WhereT(predicate).OrderByT(orderByPredicate).SelectT(toEntitySelector)
}

func (r *timeStatusDummyRepository) GetLatest() *entity.TimeStatus {
	predicate := getWhereDayPredicate(util.GetNowDateTime())
	if l, ok := linq.From(r.data).WhereT(predicate).OrderByT(orderByPredicate).Last().(datamodel.TimeStatus); ok {
		p := l.ToEntity()
		return &p
	} else {
		return nil
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
