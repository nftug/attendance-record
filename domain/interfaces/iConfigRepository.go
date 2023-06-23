package interfaces

import "attendance-record/domain/config"

type IConfigRepository interface {
	LoadConfig() (*config.Config, error)
	SaveConfig(cfg config.Config) error
}
