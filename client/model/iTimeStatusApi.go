package model

import "attendance-record/domain/dto"

type ITimeStatusApi interface {
	ToggleWork() dto.CurrentTimeStatusDto
	ToggleRest() dto.CurrentTimeStatusDto
	GetCurrentStatus() dto.CurrentTimeStatusDto
}
