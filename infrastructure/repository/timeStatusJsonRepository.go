package repository

import (
	"attendance-record/domain/entity"
	"attendance-record/domain/enum"
	"attendance-record/domain/interfaces"
	"attendance-record/infrastructure/datamodel"
	"attendance-record/shared/util"
	"fmt"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/sonyarouje/simdb"
)

func NewWorkJsonRepository() interfaces.IWorkRepository {
	return newTimeStatusJsonRepository(enum.Work)
}

func NewRestJsonRepository() interfaces.IRestRepository {
	return newTimeStatusJsonRepository(enum.Rest)
}

type timeStatusJsonRepository struct {
	driver *simdb.Driver
}

func newTimeStatusJsonRepository(recordType enum.TimeStatusType) *timeStatusJsonRepository {
	var fn string
	switch recordType {
	case enum.Work:
		fn = "work_record"
	case enum.Rest:
		fn = "rest_record"
	}

	driver, err := simdb.New(fn)
	if err != nil {
		panic(err)
	}
	return &timeStatusJsonRepository{driver: driver}
}

func (r *timeStatusJsonRepository) Create(item entity.TimeStatus) {
	d := datamodel.NewTimeStatusFromEntity(item)
	r.driver.Insert(d)
}

func (r *timeStatusJsonRepository) Update(item entity.TimeStatus) {
	fmt.Println("Update")
	updated := datamodel.NewTimeStatusFromEntity(item)
	r.driver.Update(updated)
}

func (r *timeStatusJsonRepository) QueryByDate(dt time.Time) linq.Query {
	var results []datamodel.TimeStatus
	today, _ := getDayPair(dt)
	err := r.driver.Open(datamodel.TimeStatus{}).
		Where("StartTime", ">=", today).
		Get().AsEntity(&results)
	if err != nil {
		panic(err)
	}
	return linq.From(results).SelectT(toEntitySelector)
}

func (r *timeStatusJsonRepository) GetLatest() *entity.TimeStatus {
	var results []datamodel.TimeStatus
	today, _ := getDayPair(util.GetNowDateTime())
	err := r.driver.Open(datamodel.TimeStatus{}).
		Where("StartTime", ">=", today).
		Get().AsEntity(&results)
	if err != nil {
		panic(err)
	}

	if l, ok := linq.From(results).
		WhereT(getWhereDayPredicate(util.GetNowDateTime())).
		OrderByT(orderByPredicate).
		Last().(datamodel.TimeStatus); ok {
		p := l.ToEntity()
		return &p
	} else {
		return nil
	}
}
