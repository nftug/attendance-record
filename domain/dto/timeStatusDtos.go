package dto

import (
	"attendance-record/domain/enum"
	"time"

	"github.com/google/uuid"
)

type TimeStatusDto struct {
	Id        uuid.UUID
	Type      enum.TimeStatusType
	StartedOn time.Time
	EndedOn   time.Time
	TotalTime time.Duration
}

type CurrentTimeStatusItemDto struct {
	IsToggleEnabled bool
	IsActive        bool
	TotalTime       time.Duration
	StartedOn       time.Time
	EndedOn         time.Time
}

type CurrentTimeStatusDto struct {
	Work CurrentTimeStatusItemDto
	Rest CurrentTimeStatusItemDto
}
