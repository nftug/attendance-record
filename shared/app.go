package shared

import (
	"attendance-record/domain/config"
	"attendance-record/usecase"
)

var app *App

type App struct {
	Config            *config.Config
	ConfigRepository  config.IConfigRepository
	TimeStatusUseCase *usecase.TimeStatusUseCase
}

func NewAppSingleton(tsUseCase *usecase.TimeStatusUseCase, cfgRepo config.IConfigRepository) *App {
	if app == nil {
		cfg, _ := cfgRepo.LoadConfig()
		app = &App{
			TimeStatusUseCase: tsUseCase,
			Config:            cfg,
			ConfigRepository:  cfgRepo,
		}
	}
	return app
}
