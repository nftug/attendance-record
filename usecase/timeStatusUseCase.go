package usecase

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/entity"
	"attendance-record/domain/enum"
	"attendance-record/domain/interfaces"
	"attendance-record/domain/service"
	"time"

	"github.com/google/uuid"
)

type TimeStatusUseCase struct {
	service *service.TimeStatusService
	repo    *interfaces.TimeStatusRepositorySet
}

func NewTimeStatusUseCase(service *service.TimeStatusService, repo *interfaces.TimeStatusRepositorySet) *TimeStatusUseCase {
	return &TimeStatusUseCase{service, repo}
}

func (u *TimeStatusUseCase) ToggleWork() error {
	return u.service.ToggleState(enum.Work)
}

func (u *TimeStatusUseCase) ToggleRest() error {
	return u.service.ToggleState(enum.Rest)
}

func (u *TimeStatusUseCase) GetCurrent() (dto.CurrentTimeStatusDto, error) {
	return u.service.GetCurrent()
}

func (u *TimeStatusUseCase) FindByMonth(year int, month time.Month) ([]dto.TimeStatusDto, error) {
	return u.service.FindByMonth(year, month)
}

func (u *TimeStatusUseCase) Create(t enum.TimeStatusType, cmd dto.TimeStatusCommandDto) error {
	repo := u.repo.Get(t)
	item, err := entity.NewTimeStatus(cmd)
	if err != nil {
		return err
	}
	repo.Create(*item)
	return nil
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
