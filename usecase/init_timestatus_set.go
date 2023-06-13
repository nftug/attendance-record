package usecase

import "domain/entity"

func InitTimeStatusSet() *entity.TimeStatusSet {
	return entity.NewTimeStatusSet()
}
