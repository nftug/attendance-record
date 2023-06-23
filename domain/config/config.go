package config

import "time"

type Config struct {
	WorkHours int `json:"workHours"`
}

func (c *Config) Overtime(workHrs time.Duration) time.Duration {
	return workHrs - time.Duration(int64(c.WorkHours))*time.Hour
}
