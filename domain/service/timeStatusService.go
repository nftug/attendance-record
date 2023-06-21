package service

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/entity"
	"attendance-record/domain/enum"
	"attendance-record/domain/interfaces"
	"attendance-record/shared/util"
	"errors"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/google/uuid"
)

type TimeStatusService struct {
	workRepository interfaces.IWorkRepository
	restRepository interfaces.IRestRepository
}

func NewTimeStatusService(wr interfaces.IWorkRepository, rr interfaces.IRestRepository) *TimeStatusService {
	return &TimeStatusService{wr, rr}
}

func (tss *TimeStatusService) ToggleState(t enum.TimeStatusType) error {
	if t == enum.Work && tss.isActiveByRepository(enum.Rest) {
		return errors.New("cannot toggle work status")
	} else if t == enum.Rest && !tss.isActiveByRepository(enum.Work) {
		return errors.New("cannot toggle rest status")
	}

	repo := tss.getRepository(t)
	if item := repo.GetLatest(); item != nil && item.IsActive() {
		item.End()
		repo.Update(*item)
	} else {
		if t == enum.Work &&
			item != nil &&
			util.GetDate(item.EndTime) == util.GetDate(time.Now()) {
			return errors.New("no operation is allowed after leaving work within the same day")
		}
		repo.Create(entity.NewTimeStatus())
	}

	return nil
}

func (tss *TimeStatusService) GetCurrent() dto.CurrentTimeStatusDto {
	var workStartedOn, workEndedOn, restStartedOn, restEndedOn time.Time

	now := util.GetNowDateTime()
	queryWork := tss.workRepository.QueryByDate(now)
	queryRest := tss.restRepository.QueryByDate(now)
	selTotal := func(x entity.TimeStatus) int64 { return int64(x.TotalTime(now)) }

	workTotal := queryWork.SelectT(selTotal).SumInts()
	if wf, ok := queryWork.First().(entity.TimeStatus); ok {
		workStartedOn = wf.StartTime
	}
	if wl, ok := queryWork.Last().(entity.TimeStatus); ok && !wl.IsActive() {
		workEndedOn = wl.EndTime
	}

	restTotal := queryRest.SelectT(selTotal).SumInts()
	if rl, ok := queryRest.Last().(entity.TimeStatus); ok {
		restStartedOn = rl.StartTime
		restEndedOn = rl.EndTime
	}

	return dto.CurrentTimeStatusDto{
		Work: dto.CurrentTimeStatusItemDto{
			IsToggleEnabled: isWorkToggleEnabled(queryWork, queryRest),
			IsActive:        isActiveByQuery(queryWork),
			TotalTime:       time.Duration(workTotal - restTotal),
			StartedOn:       workStartedOn,
			EndedOn:         workEndedOn,
		},
		Rest: dto.CurrentTimeStatusItemDto{
			IsToggleEnabled: isActiveByQuery(queryWork),
			IsActive:        isActiveByQuery(queryRest),
			TotalTime:       time.Duration(restTotal),
			StartedOn:       restStartedOn,
			EndedOn:         restEndedOn,
		},
	}
}

func (tss *TimeStatusService) GetAll() (results []dto.TimeStatusDto) {
	workAll := tss.workRepository.GetAll()
	restAll := tss.restRepository.GetAll()
	now := util.GetNowDateTime()

	type RestDuration struct {
		// workId uuid.UUID
		date     time.Time
		duration time.Duration
	}

	restDurations := restAll.
		GroupByT(
			func(x entity.TimeStatus) time.Time {
				return util.GetDate(x.StartTime)
				// return workAll.
				//	FirstWithT(func(w entity.TimeStatus) bool { return x.StartTime.After(w.StartTime) }).(entity.TimeStatus).Id
			},
			func(x entity.TimeStatus) int64 { return int64(x.TotalTime(now)) },
		).
		SelectT(func(x linq.Group) RestDuration {
			return RestDuration{
				date: x.Key.(time.Time),
				// workId:   x.Key.(uuid.UUID),
				duration: time.Duration(linq.From(x.Group).SumInts()),
			}
		})

	workDtos := workAll.GroupJoinT(
		restDurations,
		func(x entity.TimeStatus) time.Time { return util.GetDate(x.StartTime) },
		func(x RestDuration) time.Time { return x.date },
		func(o entity.TimeStatus, i []RestDuration) dto.TimeStatusDto {
			var totalTime time.Duration
			if len(i) > 0 {
				totalTime = i[0].duration
			}
			return dto.TimeStatusDto{
				Id:        o.Id,
				Type:      enum.Work,
				StartedOn: o.StartTime,
				EndedOn:   o.EndTime,
				TotalTime: o.TotalTime(now) - totalTime,
			}
		},
	)

	restDtos := restAll.SelectT(func(x entity.TimeStatus) dto.TimeStatusDto {
		return dto.TimeStatusDto{
			Id:        x.Id,
			Type:      enum.Rest,
			StartedOn: x.StartTime,
			EndedOn:   x.EndTime,
			TotalTime: x.TotalTime(now),
		}
	})

	workDtos.
		Concat(restDtos).
		OrderByT(func(x dto.TimeStatusDto) int64 { return x.StartedOn.Unix() }).
		ToSlice(&results)
	return
}

func (tss *TimeStatusService) Delete(t enum.TimeStatusType, id uuid.UUID) error {
	return tss.getRepository(t).Delete(id)
}

func (tss *TimeStatusService) Update(t enum.TimeStatusType, id uuid.UUID, cmd dto.TimeStatusCommandDto) error {
	repo := tss.getRepository(t)
	item, err := repo.Get(id)
	if err != nil {
		return err
	}
	if err := item.Edit(cmd); err != nil {
		return err
	}
	repo.Update(*item)
	return nil
}

func (tss *TimeStatusService) isActiveByRepository(t enum.TimeStatusType) bool {
	l := tss.getRepository(t).GetLatest()
	return l != nil && l.IsActive()
}

func isActiveByQuery(q linq.Query) bool {
	l, ok := q.Last().(entity.TimeStatus)
	return ok && l.IsActive()
}

func isWorkToggleEnabled(work linq.Query, rest linq.Query) bool {
	wl, ok := work.Last().(entity.TimeStatus)
	isWorkActive := ok &&
		wl.IsActive() ||
		util.GetDate(wl.EndTime) != util.GetDate(time.Now())
	return !isActiveByQuery(rest) && isWorkActive
}

func (tss *TimeStatusService) getRepository(t enum.TimeStatusType) interfaces.ITimeStatusRepository {
	if t == enum.Work {
		return tss.workRepository
	} else if t == enum.Rest {
		return tss.restRepository
	}
	return nil
}
