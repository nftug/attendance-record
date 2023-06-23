package usecase

import (
	"attendance-record/domain/config"
	"attendance-record/domain/interfaces"
)

type ConfigUseCase struct {
	repo interfaces.IConfigRepository
}

func NewConfigUseCase(repo interfaces.IConfigRepository) *ConfigUseCase {
	return &ConfigUseCase{repo}
}

func (u *ConfigUseCase) LoadConfig() (*config.Config, error) {
	return u.repo.LoadConfig()
}

func (u *ConfigUseCase) SaveConfig(cfg config.Config) error {
	return u.repo.SaveConfig(cfg)
}
