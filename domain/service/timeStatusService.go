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
)

type TimeStatusService struct {
	repo *interfaces.TimeStatusRepositorySet
}

func NewTimeStatusService(r *interfaces.TimeStatusRepositorySet) *TimeStatusService {
	return &TimeStatusService{r}
}

func (tss *TimeStatusService) ToggleState(t enum.TimeStatusType) error {
	if t == enum.Work && tss.isActiveByRepository(enum.Rest) {
		return errors.New("cannot toggle work status")
	} else if t == enum.Rest && !tss.isActiveByRepository(enum.Work) {
		return errors.New("cannot toggle rest status")
	}

	repo := tss.repo.Get(t)
	if item, err := repo.GetLatest(); item != nil && item.IsActive() {
		item.End()
		repo.Update(*item)
	} else if err != nil {
		return err
	} else {
		if t == enum.Work &&
			item != nil &&
			util.GetDate(item.Record.EndTime) == util.GetDate(time.Now()) {
			return errors.New("no operation is allowed after leaving work within the same day")
		}
		repo.Create(entity.NewTimeStatusAsNow())
	}

	return nil
}

func (tss *TimeStatusService) GetCurrent() (dto.CurrentTimeStatusDto, error) {
	var workStartedOn, workEndedOn, restStartedOn, restEndedOn time.Time

	now := util.GetNowDateTime()
	today := util.GetDate(now)
	tomorrow := today.AddDate(0, 0, 1)
	queryWork, err := tss.repo.Get(enum.Work).FindByDate(today, tomorrow)
	if err != nil {
		return *new(dto.CurrentTimeStatusDto), err
	}
	queryRest, err := tss.repo.Get(enum.Rest).FindByDate(today, tomorrow)
	if err != nil {
		return *new(dto.CurrentTimeStatusDto), err
	}
	selTotal := func(x entity.TimeStatus) int64 { return int64(x.TotalTime(now)) }

	workTotal := queryWork.SelectT(selTotal).SumInts()
	if wf, ok := queryWork.First().(entity.TimeStatus); ok {
		workStartedOn = wf.Record.StartTime
	}
	if wl, ok := queryWork.Last().(entity.TimeStatus); ok && !wl.IsActive() {
		workEndedOn = wl.Record.EndTime
	}

	restTotal := queryRest.SelectT(selTotal).SumInts()
	if rl, ok := queryRest.Last().(entity.TimeStatus); ok {
		restStartedOn = rl.Record.StartTime
		restEndedOn = rl.Record.EndTime
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
	}, nil
}

func (tss *TimeStatusService) FindByMonth(year int, month time.Month) (results []dto.TimeStatusDto, err error) {
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 1, 0)

	workAll, err := tss.repo.Get(enum.Work).FindByDate(startDate, endDate)
	if err != nil {
		return nil, err
	}
	restAll, err := tss.repo.Get(enum.Rest).FindByDate(startDate, endDate)
	if err != nil {
		return nil, err
	}
	now := util.GetNowDateTime()

	type RestDuration struct {
		date     time.Time
		duration time.Duration
	}

	restDurations := restAll.
		GroupByT(
			func(x entity.TimeStatus) time.Time { return util.GetDate(x.Record.StartTime) },
			func(x entity.TimeStatus) int64 { return int64(x.TotalTime(now)) },
		).
		SelectT(func(x linq.Group) RestDuration {
			return RestDuration{
				date:     x.Key.(time.Time),
				duration: time.Duration(linq.From(x.Group).SumInts()),
			}
		})

	workDtos := workAll.GroupJoinT(
		restDurations,
		func(x entity.TimeStatus) time.Time { return util.GetDate(x.Record.StartTime) },
		func(x RestDuration) time.Time { return x.date },
		func(o entity.TimeStatus, i []RestDuration) dto.TimeStatusDto {
			var totalTime time.Duration
			if len(i) > 0 {
				totalTime = i[0].duration
			}
			return dto.TimeStatusDto{
				Id:        o.Id,
				Type:      enum.Work,
				StartedOn: o.Record.StartTime,
				EndedOn:   o.Record.EndTime,
				TotalTime: o.TotalTime(now) - totalTime,
			}
		},
	)

	restDtos := restAll.SelectT(func(x entity.TimeStatus) dto.TimeStatusDto {
		return dto.TimeStatusDto{
			Id:        x.Id,
			Type:      enum.Rest,
			StartedOn: x.Record.StartTime,
			EndedOn:   x.Record.EndTime,
			TotalTime: x.TotalTime(now),
		}
	})

	workDtos.
		Concat(restDtos).
		OrderByT(func(x dto.TimeStatusDto) int64 { return x.StartedOn.Unix() }).
		ToSlice(&results)
	return
}

func (tss *TimeStatusService) isActiveByRepository(t enum.TimeStatusType) bool {
	l, _ := tss.repo.Get(t).GetLatest()
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
		util.GetDate(wl.Record.EndTime) != util.GetDate(time.Now())
	return !isActiveByQuery(rest) && isWorkActive
}
