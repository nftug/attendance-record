package util

import "time"

func GetDateTime(dt time.Time) time.Time {
	return time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), dt.Second(), 0, time.Local)
}

func GetNowDateTime() time.Time {
	return GetDateTime(time.Now())
}
