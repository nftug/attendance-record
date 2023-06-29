package config

import "time"

type RestAlarm struct {
	IsEnabled bool `json:"isEnabled"`
	Hours     int  `json:"hours"`
	Minutes   int  `json:"minutes"`
}

func (r *RestAlarm) ShouldInvoke(c *Config, workHrs time.Duration, restHrs time.Duration) bool {
	if !r.IsEnabled || workHrs == 0 || restHrs > 0 {
		return false
	}
	duration := time.Duration(r.Hours*int(time.Hour) + r.Minutes*int(time.Minute))
	return c.Overtime(workHrs) >= duration
}
