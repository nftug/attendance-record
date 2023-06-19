package shared

import "attendance-record/usecase"

type App struct {
	TimeStatusUseCase *usecase.TimeStatusUseCase
}

func NewApp(tsUseCase *usecase.TimeStatusUseCase) *App {
	return &App{TimeStatusUseCase: tsUseCase}
}
