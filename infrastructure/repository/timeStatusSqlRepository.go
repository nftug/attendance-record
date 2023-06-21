package repository

import (
	"attendance-record/domain/entity"
	"attendance-record/domain/enum"
	"attendance-record/domain/interfaces"
	"attendance-record/infrastructure/datamodel"
	"attendance-record/shared/util"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewWorkSqlRepository(db *gorm.DB) interfaces.IWorkRepository {
	return &timeStatusSqlRepository{db: getDbModel(db, enum.Work)}
}

func NewRestSqlRepository(db *gorm.DB) interfaces.IRestRepository {
	return &timeStatusSqlRepository{db: getDbModel(db, enum.Rest)}
}

type timeStatusSqlRepository struct {
	db *gorm.DB
}

func (r *timeStatusSqlRepository) Create(item entity.TimeStatus) {
	d := datamodel.NewTimeStatusFromEntity(item)
	r.db.Create(&d)
}

func (r *timeStatusSqlRepository) Update(item entity.TimeStatus) {
	updated := datamodel.NewTimeStatusFromEntity(item)
	e := r.db.First(&datamodel.TimeStatus{}, item.Id)
	e.Updates(&updated)

	if updated.EndTime == *new(time.Time) {
		e.Select("end_time").Updates(updated)
	}
}

func (r *timeStatusSqlRepository) QueryByDate(dt time.Time) linq.Query {
	var entities []datamodel.TimeStatus
	today, tomorrow := getDayPair(dt)
	r.db.
		Where("start_time BETWEEN ? AND ?", today, tomorrow).
		Order("start_time").
		Find(&entities)
	return linq.From(entities).SelectT(toEntitySelector)
}

func (r *timeStatusSqlRepository) GetAll() linq.Query {
	var entities []datamodel.TimeStatus
	r.db.Order("start_time").Find(&entities)
	return linq.From(entities).SelectT(toEntitySelector)
}

func (r *timeStatusSqlRepository) GetLatest() *entity.TimeStatus {
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

func (r *timeStatusSqlRepository) Delete(id uuid.UUID) error {
	ret := r.db.Delete(&datamodel.TimeStatus{}, id)
	return ret.Error
}

func (r *timeStatusSqlRepository) Get(id uuid.UUID) (*entity.TimeStatus, error) {
	var e datamodel.TimeStatus
	ret := r.db.First(&e, id)
	return &entity.TimeStatus{
		Id:        e.Id,
		StartTime: e.StartTime,
		EndTime:   e.EndTime,
	}, ret.Error
}

func getDbModel(db *gorm.DB, recordType enum.TimeStatusType) *gorm.DB {
	if recordType == enum.Work {
		return db.Model(datamodel.WorkTimeStatus{}).Session(&gorm.Session{})
	} else {
		return db.Model(datamodel.RestTimeStatus{}).Session(&gorm.Session{})
	}
}
