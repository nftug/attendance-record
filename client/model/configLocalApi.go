package model

import "attendance-record/domain/config"

type configLocalApi struct {
	repository config.IConfigRepository
}

var Config *config.Config

func NewConfigLocalApi(r config.IConfigRepository) IConfigApi {
	api := &configLocalApi{r}
	Config, _ = api.LoadConfig()
	return api
}

func (api *configLocalApi) LoadConfig() (*config.Config, error) {
	cfg, err := api.repository.LoadConfig()
	if err != nil {
		return nil, err
	}
	Config = cfg
	return cfg, err
}

func (api *configLocalApi) SaveConfig(cfg config.Config) error {
	if err := api.repository.SaveConfig(cfg); err != nil {
		return err
	}
	Config = &cfg
	return nil
}
