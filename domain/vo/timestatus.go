package vo

import (
	"fmt"
	"time"
)

type TimeStatus struct {
	IsToggleEnabled bool
	IsActive        bool
	StartTime       time.Time
	EndTime         time.Time
	TotalTime       time.Duration
}

func (ts *TimeStatus) ToggleActive() {
	if !ts.IsToggleEnabled {
		return
	}

	ts.IsActive = !ts.IsActive

	if ts.IsActive {
		ts.StartTime = time.Now()
		ts.EndTime = time.Time{}
	} else {
		ts.EndTime = time.Now()
		ts.TotalTime += ts.EndTime.Sub(ts.StartTime)
	}
}

func (ts *TimeStatus) OnTickTimer() {
	if !ts.IsActive || !ts.IsToggleEnabled {
		fmt.Println("Not active")
		return
	}
	ts.TotalTime += time.Duration(1) * time.Second
}
