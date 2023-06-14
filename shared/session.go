package shared

import (
	"domain/service"
	"infrastructure/repository"
)

type Session struct {
	TimeStatusService *service.TimeStatusService
}

func NewSession() *Session {
	wr := repository.NewTimeStatusDummyRepository()
	rr := repository.NewTimeStatusDummyRepository()

	return &Session{TimeStatusService: service.NewTimeStatusService(wr, rr)}
}
