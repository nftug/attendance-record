package config

import "time"

type WorkAlarm struct {
	IsEnabled     bool `json:"isEnabled"`
	BeforeMinutes int  `json:"beforeMinutes"`
}

func (w *WorkAlarm) ShouldInvoke(c *Config, workHrs time.Duration) bool {
	if !w.IsEnabled || workHrs == 0 {
		return false
	}
	return c.Overtime(workHrs) >= time.Duration(w.BeforeMinutes*int(time.Minute))*-1
}
