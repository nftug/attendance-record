package config

import "time"

type Config struct {
	WorkHours int       `json:"workHours"`
	WorkAlarm WorkAlarm `json:"workAlarm"`
	RestAlarm RestAlarm `json:"restAlarm"`
}

var DefaultConfig = Config{
	WorkHours: 8,
	WorkAlarm: WorkAlarm{
		IsEnabled:     true,
		BeforeMinutes: 15,
		SnoozeMinutes: 5,
	},
	RestAlarm: RestAlarm{
		IsEnabled:     false,
		Hours:         4,
		Minutes:       0,
		SnoozeMinutes: 5,
	},
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
