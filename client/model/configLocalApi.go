package model

import (
	"attendance-record/domain/config"
	"attendance-record/shared"
	"attendance-record/usecase"
)

type configLocalApi struct {
	usecase *usecase.ConfigUseCase
}

var Config *config.Config

func NewConfigLocalApi(a *shared.App) IConfigApi {
	api := &configLocalApi{a.ConfigUseCase}
	Config = a.Config
	return api
}

func (api *configLocalApi) LoadConfig() (*config.Config, error) {
	cfg, err := api.usecase.LoadConfig()
	if err != nil {
		return nil, err
	}
	Config = cfg
	return cfg, err
}

func (api *configLocalApi) SaveConfig(cfg config.Config) error {
	if err := api.usecase.SaveConfig(cfg); err != nil {
		return err
	}
	Config = &cfg
	return nil
}
