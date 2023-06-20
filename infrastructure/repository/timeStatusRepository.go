package repository

import (
	"attendance-record/domain/entity"
	"attendance-record/domain/enum"
	"attendance-record/domain/interfaces"
	"attendance-record/infrastructure/datamodel"
	"attendance-record/shared/util"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"gorm.io/gorm"
)

func NewWorkRepository(db *gorm.DB) interfaces.IWorkRepository {
	return &timeStatusRepository{db: getDbModel(db, enum.Work)}
}

func NewRestRepository(db *gorm.DB) interfaces.IRestRepository {
	return &timeStatusRepository{db: getDbModel(db, enum.Rest)}
}

type timeStatusRepository struct {
	db *gorm.DB
}

func (r *timeStatusRepository) Create(item entity.TimeStatus) {
	d := datamodel.NewTimeStatusFromEntity(item)
	r.db.Create(&d)
}

func (r *timeStatusRepository) Update(item entity.TimeStatus) {
	updated := datamodel.NewTimeStatusFromEntity(item)
	r.db.First(&datamodel.TimeStatus{}, item.Id).Updates(&updated)
}

func (r *timeStatusRepository) QueryByDate(dt time.Time) linq.Query {
	var results []datamodel.TimeStatus
	today, tomorrow := getDayPair(dt)
	r.db.
		Where("start_time BETWEEN ? AND ?", today, tomorrow).
		Order("start_time").
		Find(&results)
	return linq.From(results).SelectT(toEntitySelector)
}

func (r *timeStatusRepository) GetLatest() *entity.TimeStatus {
	var entity datamodel.TimeStatus
	today, tomorrow := getDayPair(util.GetNowDateTime())
	r.db.
		Where("start_time BETWEEN ? AND ?", today, tomorrow).
		Order("start_time DESC").
		FirstOrInit(&entity)

	if entity != *new(datamodel.TimeStatus) {
		p := entity.ToEntity()
		return &p
	} else {
		return nil
	}
}

func getDbModel(db *gorm.DB, recordType enum.TimeStatusType) *gorm.DB {
	if recordType == enum.Work {
		return db.Model(datamodel.WorkTimeStatus{}).Session(&gorm.Session{})
	} else {
		return db.Model(datamodel.RestTimeStatus{}).Session(&gorm.Session{})
	}
}
