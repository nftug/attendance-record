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

func (r *timeStatusSqlRepository) Create(item entity.TimeStatus) error {
	d := datamodel.NewTimeStatusFromEntity(item)
	return r.db.Create(&d).Error
}

func (r *timeStatusSqlRepository) Update(item entity.TimeStatus) error {
	updated := datamodel.NewTimeStatusFromEntity(item)
	e := r.db.First(&datamodel.TimeStatus{}, item.Id)
	if err := e.Updates(&updated).Error; err != nil {
		return err
	}
	if updated.EndTime == *new(time.Time) {
		return e.Select("end_time").Updates(updated).Error
	}

	return nil
}

func (r *timeStatusSqlRepository) FindByDate(start time.Time, end time.Time) (linq.Query, error) {
	var entities []datamodel.TimeStatus
	ctx := r.db.
		Where("start_time BETWEEN ? AND ?", util.GetDate(start), util.GetDate(end)).
		Order("start_time").
		Find(&entities)
	return linq.From(entities).SelectT(toEntitySelector), ctx.Error
}

func (r *timeStatusSqlRepository) GetLatest() (*entity.TimeStatus, error) {
	var entity datamodel.TimeStatus
	today, tomorrow := getDayPair(util.GetNowDateTime())
	ctx := r.db.
		Where("start_time BETWEEN ? AND ?", today, tomorrow).
		Order("start_time DESC").
		FirstOrInit(&entity)

	if entity != *new(datamodel.TimeStatus) {
		p := entity.ToEntity()
		return &p, nil
	} else {
		return nil, ctx.Error
	}
}

func (r *timeStatusSqlRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&datamodel.TimeStatus{}, id).Error
}

func (r *timeStatusSqlRepository) Get(id uuid.UUID) (*entity.TimeStatus, error) {
	var e datamodel.TimeStatus
	ctx := r.db.First(&e, id)
	p := e.ToEntity()
	return &p, ctx.Error
}

func getDbModel(db *gorm.DB, recordType enum.TimeStatusType) *gorm.DB {
	if recordType == enum.Work {
		return db.Model(datamodel.WorkTimeStatus{}).Session(&gorm.Session{})
	} else {
		return db.Model(datamodel.RestTimeStatus{}).Session(&gorm.Session{})
	}
}
