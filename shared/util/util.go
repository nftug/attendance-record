package util

import "time"

func GetDateTime(dt time.Time) time.Time {
	return time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), dt.Second(), 0, time.Local)
}

func GetDate(dt time.Time) time.Time {
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, time.Local)
}

func GetNowDateTime() time.Time {
	return GetDateTime(time.Now())
}

func FormatDateTime(dt time.Time) string {
	if dt == *new(time.Time) {
		return ""
	} else {
		return dt.Format("15:04")
	}
}

func SetHourAndMinute(origin time.Time, sub time.Time) time.Time {
	origin = GetDate(origin)
	return time.Date(origin.Year(), origin.Month(), origin.Day(), sub.Hour(), sub.Minute(), sub.Second(), 0, time.Local)
}
