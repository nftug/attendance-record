package repository

import "attendance-record/domain/interfaces"

func NewWorkDummyRepository() interfaces.WorkRepository {
	return &timeStatusDummyRepository{}
}

func NewRestDummyRepository() interfaces.RestRepository {
	return &timeStatusDummyRepository{}
}
