package shared

import (
	"attendance-record/usecase"
)

var app *App

type App struct {
	ConfigUseCase     *usecase.ConfigUseCase
	TimeStatusUseCase *usecase.TimeStatusUseCase
}

func NewAppSingleton(tsUseCase *usecase.TimeStatusUseCase, cfgUseCase *usecase.ConfigUseCase) *App {
	if app == nil {
		app = &App{
			TimeStatusUseCase: tsUseCase,
			ConfigUseCase:     cfgUseCase,
		}
	}
	return app
}
