package shared

import "domain/entity"

type Session struct {
	TimeStatusSet *entity.TimeStatusSet
}

func NewSession() *Session {
	return &Session{TimeStatusSet: entity.NewTimeStatusSet()}
}
