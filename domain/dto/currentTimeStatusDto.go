package dto

import "time"

type TimeStatusItemDto struct {
	IsToggleEnabled bool
	IsActive        bool
	TotalTime       time.Duration
	StartedOn       time.Time
	EndedOn         time.Time
}

type CurrentTimeStatusDto struct {
	Work TimeStatusItemDto
	Rest TimeStatusItemDto
}
