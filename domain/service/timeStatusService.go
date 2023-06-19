package service

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/entity"
	"attendance-record/domain/enum"
	"attendance-record/domain/interfaces"
	"time"

	"github.com/ahmetb/go-linq/v3"
)

type TimeStatusService struct {
	workRepository interfaces.IWorkRepository
	restRepository interfaces.IRestRepository
}

func NewTimeStatusService(wr interfaces.IWorkRepository, rr interfaces.IRestRepository) *TimeStatusService {
	return &TimeStatusService{wr, rr}
}

func (tss *TimeStatusService) ToggleState(t enum.TimeStatusType) {
	if t == enum.Work && tss.isActiveByRepository(enum.Rest) {
		return
	} else if t == enum.Rest && !tss.isActiveByRepository(enum.Work) {
		return
	}

	repo := tss.getRepository(t)
	if item := repo.GetLatest(); item != nil && item.IsActive() {
		item.End()
		repo.Update(*item)
	} else {
		repo.Create(entity.NewTimeStatus())
	}
}

func (tss *TimeStatusService) GetCurrent() dto.CurrentTimeStatusDto {
	var workStartedOn, workEndedOn, restStartedOn, restEndedOn time.Time

	now := time.Now()
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
		Work: dto.TimeStatusItemDto{
			IsToggleEnabled: !isActiveByQuery(queryRest),
			IsActive:        isActiveByQuery(queryWork),
			TotalTime:       time.Duration(workTotal - restTotal),
			StartedOn:       workStartedOn,
			EndedOn:         workEndedOn,
		},
		Rest: dto.TimeStatusItemDto{
			IsToggleEnabled: isActiveByQuery(queryWork),
			IsActive:        isActiveByQuery(queryRest),
			TotalTime:       time.Duration(restTotal),
			StartedOn:       restStartedOn,
			EndedOn:         restEndedOn,
		},
	}
}

func (tss *TimeStatusService) isActiveByRepository(t enum.TimeStatusType) bool {
	l := tss.getRepository(t).GetLatest()
	return l != nil && l.IsActive()
}

func isActiveByQuery(q linq.Query) bool {
	l, ok := q.Last().(entity.TimeStatus)
	return ok && l.IsActive()
}

func (tss *TimeStatusService) getRepository(t enum.TimeStatusType) interfaces.ITimeStatusRepository {
	if t == enum.Work {
		return tss.workRepository
	} else if t == enum.Rest {
		return tss.restRepository
	}
	return nil
}
