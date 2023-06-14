package shared

import (
	"attendance-record/domain/service"
	"attendance-record/infrastructure/repository"
)

type Session struct {
	TimeStatusService *service.TimeStatusService
}

func NewSession() *Session {
	wr := repository.NewTimeStatusDummyRepository()
	rr := repository.NewTimeStatusDummyRepository()

	return &Session{TimeStatusService: service.NewTimeStatusService(wr, rr)}
}
