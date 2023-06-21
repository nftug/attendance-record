package model

import (
	"attendance-record/domain/dto"
	"attendance-record/domain/enum"

	"github.com/google/uuid"
)

type ITimeStatusApi interface {
	ToggleWork() dto.CurrentTimeStatusDto
	ToggleRest() dto.CurrentTimeStatusDto
	GetCurrentStatus() dto.CurrentTimeStatusDto
	GetAll() []dto.TimeStatusDto
	Delete(t enum.TimeStatusType, id uuid.UUID) error
	Update(t enum.TimeStatusType, id uuid.UUID, cmd dto.TimeStatusCommandDto) error
}
