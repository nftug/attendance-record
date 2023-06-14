package repository

import (
	"attendance-record/domain/entity"
	"attendance-record/domain/interfaces"
	"log"
	"time"

	"github.com/ahmetb/go-linq/v3"
)

type timeStatusDummyRepository struct {
	data []entity.TimeStatus
}

func NewTimeStatusDummyRepository() interfaces.TimeStatusRepository {
	return &timeStatusDummyRepository{}
}

func (r *timeStatusDummyRepository) Create(item entity.TimeStatus) {
	r.data = append(r.data, item)
}

func (r *timeStatusDummyRepository) Update(item entity.TimeStatus) {
	idx := linq.From(r.data).IndexOfT(func(x entity.TimeStatus) bool { return x.Id == item.Id })
	if idx == -1 {
		log.Fatal("The item with specified id cannot be found.")
	}
	r.data[idx] = item
}

func (r *timeStatusDummyRepository) QueryByDate(dt time.Time) linq.OrderedQuery {
	today := time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, time.Local)
	return linq.From(r.data).WhereT(func(x entity.TimeStatus) bool {
		return x.StartTime.After(today)
	}).OrderByT(func(x entity.TimeStatus) int64 {
		return x.StartTime.Unix()
	})
}
