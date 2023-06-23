package config

import "time"

type Config struct {
	WorkHours int `json:"workHours"`
}

func (c *Config) OverTime(workHrs time.Duration) time.Duration {
	return workHrs - time.Duration(int64(c.WorkHours))*time.Hour
}

type IConfigRepository interface {
	LoadConfig() (*Config, error)
	SaveConfig(cfg Config) error
}
