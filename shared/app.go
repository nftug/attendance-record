package shared

import (
	"attendance-record/domain/config"
	"attendance-record/usecase"
)

var app *App

type App struct {
	Config            *config.Config
	ConfigUseCase     *usecase.ConfigUseCase
	TimeStatusUseCase *usecase.TimeStatusUseCase
}

func NewAppSingleton(tsUseCase *usecase.TimeStatusUseCase, cfgUseCase *usecase.ConfigUseCase) *App {
	if app == nil {
		cfg, _ := cfgUseCase.LoadConfig()
		app = &App{
			TimeStatusUseCase: tsUseCase,
			Config:            cfg,
			ConfigUseCase:     cfgUseCase,
		}
	}
	return app
}
