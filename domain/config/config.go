package config

import "time"

type Config struct {
	WorkHours int       `json:"workHours"`
	WorkAlarm WorkAlarm `json:"workAlarm"`
	RestAlarm RestAlarm `json:"restAlarm"`
}

func (c *Config) Overtime(workHrs time.Duration) time.Duration {
	return workHrs - time.Duration(int64(c.WorkHours))*time.Hour
}

func (c *Config) ShouldInvokeWorkAlarm(workHrs time.Duration) bool {
	return c.WorkAlarm.ShouldInvoke(c, workHrs)
}

func (c *Config) ShouldInvokeRestAlarm(workHrs time.Duration, restHrs time.Duration) bool {
	return c.RestAlarm.ShouldInvoke(c, workHrs, restHrs)
}
