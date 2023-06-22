package dto

import (
	"attendance-record/shared/util"
	"time"
)

type TimeStatusCommandDto struct {
	StartedOn time.Time
	EndedOn   time.Time
}

const DateFormat = "2006-01-02"

func NewTimeStatusCommandDto(date string, start string, end string) (TimeStatusCommandDto, error) {
	tLayout, dLayout := "15:04", DateFormat
	cmd := TimeStatusCommandDto{}

	d, err := time.Parse(dLayout, date)
	if err != nil {
		return *new(TimeStatusCommandDto), err
	}

	startedOn, err := time.Parse(tLayout, start)
	if err != nil {
		return *new(TimeStatusCommandDto), err
	}
	cmd.StartedOn = util.SetHourAndMinute(d, startedOn)

	if end != "" {
		endedOn, err := time.Parse(tLayout, end)
		if err != nil {
			return *new(TimeStatusCommandDto), err
		}
		cmd.EndedOn = util.SetHourAndMinute(d, endedOn)
	} else {
		cmd.EndedOn = *new(time.Time)
	}

	return cmd, nil
}
