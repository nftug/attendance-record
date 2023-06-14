package service

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/entity"
	"attendance-record/domain/interfaces"
	"time"
)

type TimeStatusService struct {
	workRepository interfaces.TimeStatusRepository
	restRepository interfaces.TimeStatusRepository
}

func NewTimeStatusService(wr interfaces.TimeStatusRepository, rr interfaces.TimeStatusRepository) *TimeStatusService {
	return &TimeStatusService{wr, rr}
}

func (tss *TimeStatusService) ToggleWork() {
	if tss.isResting() {
		return
	}

	l, ok := tss.workRepository.QueryByDate(time.Now()).Last().(entity.TimeStatus)
	if ok && l.IsActive() {
		// 退勤処理
		l.End()
		tss.workRepository.Update(l)
	} else {
		// 出勤処理
		tss.workRepository.Create(*entity.NewTimeStatus())
	}
}

func (tss *TimeStatusService) ToggleRest() {
	if !tss.isWorking() {
		return
	}

	l, ok := tss.restRepository.QueryByDate(time.Now()).Last().(entity.TimeStatus)
	if ok && l.IsActive() {
		// 休憩終了
		l.End()
		tss.restRepository.Update(l)
	} else {
		// 休憩開始
		tss.restRepository.Create(*entity.NewTimeStatus())
	}
}

func (tss *TimeStatusService) GetCurrent() *dto.CurrentTimeStatusDto {
	var workStartedOn, workEndedOn, restStartedOn, restEndedOn time.Time

	now := time.Now()
	queryWork := tss.workRepository.QueryByDate(now)
	queryRest := tss.restRepository.QueryByDate(now)

	workTotal := queryWork.SelectT(func(x entity.TimeStatus) int64 {
		return int64(x.TotalTime(now))
	}).SumInts()
	if wf, ok := queryWork.First().(entity.TimeStatus); ok {
		workStartedOn = wf.StartTime
	}
	if wl, ok := queryWork.Last().(entity.TimeStatus); ok && !wl.IsActive() {
		workEndedOn = wl.EndTime
	}

	restTotal := queryRest.SelectT(func(x entity.TimeStatus) int64 {
		return int64(x.TotalTime(now))
	}).SumInts()
	if rl, ok := queryRest.Last().(entity.TimeStatus); ok {
		restStartedOn = rl.StartTime
		restEndedOn = rl.EndTime
	}

	return &dto.CurrentTimeStatusDto{
		Work: dto.TimeStatusItemDto{
			IsToggleEnabled: !tss.isResting(),
			IsActive:        tss.isWorking(),
			TotalTime:       time.Duration(workTotal - restTotal),
			StartedOn:       workStartedOn,
			EndedOn:         workEndedOn,
		},
		Rest: dto.TimeStatusItemDto{
			IsToggleEnabled: tss.isWorking(),
			IsActive:        tss.isResting(),
			TotalTime:       time.Duration(restTotal),
			StartedOn:       restStartedOn,
			EndedOn:         restEndedOn,
		},
	}
}

func (tss *TimeStatusService) isResting() bool {
	// 休憩中の場合：ボタンを無効化する
	l, ok := tss.restRepository.QueryByDate(time.Now()).Last().(entity.TimeStatus)
	return ok && l.IsActive()
}

func (tss *TimeStatusService) isWorking() bool {
	// 出勤中の場合：ボタンを有効化する
	l, ok := tss.workRepository.QueryByDate(time.Now()).Last().(entity.TimeStatus)
	return ok && l.IsActive()
}
