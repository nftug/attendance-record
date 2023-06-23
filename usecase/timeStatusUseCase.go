package usecase

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/enum"
	"attendance-record/domain/interfaces"
	"attendance-record/domain/service"
	"time"

	"github.com/ahmetb/go-linq"
	"github.com/google/uuid"
)

type TimeStatusUseCase struct {
	service    *service.TimeStatusService
	repo       *interfaces.TimeStatusRepositorySet
	configRepo interfaces.IConfigRepository
}

func NewTimeStatusUseCase(
	service *service.TimeStatusService,
	repo *interfaces.TimeStatusRepositorySet,
	configRepo interfaces.IConfigRepository,
) *TimeStatusUseCase {
	return &TimeStatusUseCase{service, repo, configRepo}
}

func (u *TimeStatusUseCase) ToggleWork() error {
	return u.service.ToggleState(enum.Work)
}

func (u *TimeStatusUseCase) ToggleRest() error {
	return u.service.ToggleState(enum.Rest)
}

func (u *TimeStatusUseCase) GetCurrent() (*dto.CurrentTimeStatusDto, error) {
	return u.service.GetCurrent()
}

func (u *TimeStatusUseCase) FindByMonth(year int, month time.Month) ([]dto.TimeStatusDto, error) {
	return u.service.FindByMonth(year, month)
}

func (u *TimeStatusUseCase) Create(t enum.TimeStatusType, cmd dto.TimeStatusCommandDto) error {
	return u.service.Create(t, cmd)
}

func (u *TimeStatusUseCase) Delete(t enum.TimeStatusType, id uuid.UUID) error {
	return u.repo.Get(t).Delete(id)
}

func (u *TimeStatusUseCase) Update(t enum.TimeStatusType, id uuid.UUID, cmd dto.TimeStatusCommandDto) error {
	repo := u.repo.Get(t)
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

func (u *TimeStatusUseCase) GetOvertimeByMonth(year int, month time.Month) (*time.Duration, error) {
	data, err := u.service.FindByMonth(year, month)
	if err != nil {
		return nil, err
	}

	config, err := u.configRepo.LoadConfig()
	if err != nil {
		return nil, err
	}

	overtime := linq.From(data).
		WhereT(func(x dto.TimeStatusDto) bool { return x.Type == enum.Work }).
		SelectT(func(x dto.TimeStatusDto) int64 { return int64(config.Overtime(x.TotalTime)) }).
		SumInts()

	p := time.Duration(overtime)
	return &p, nil
}
