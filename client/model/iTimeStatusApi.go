package model

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/enum"
	"time"

	"github.com/google/uuid"
)

type ITimeStatusApi interface {
	ToggleWork() error
	ToggleRest() error
	GetCurrentStatus() (*dto.CurrentTimeStatusDto, error)
	FindByMonth(year int, month time.Month) ([]dto.TimeStatusDto, error)
	Delete(t enum.TimeStatusType, id uuid.UUID) error
	Update(t enum.TimeStatusType, id uuid.UUID, cmd dto.TimeStatusCommandDto) error
	Create(t enum.TimeStatusType, cmd dto.TimeStatusCommandDto) error
}
