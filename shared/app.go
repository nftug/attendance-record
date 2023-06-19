package shared

import "attendance-record/usecase"

var app *App

type App struct {
	TimeStatusUseCase *usecase.TimeStatusUseCase
}

func NewAppSingleton(tsUseCase *usecase.TimeStatusUseCase) *App {
	if app == nil {
		app = &App{TimeStatusUseCase: tsUseCase}
	}
	return app
}
