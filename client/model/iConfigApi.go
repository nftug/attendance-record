package model

import "attendance-record/domain/config"

type IConfigApi interface {
	LoadConfig() (*config.Config, error)
	SaveConfig(cfg config.Config) error
}
