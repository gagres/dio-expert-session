package util

import "time"

const layout = "2006-01-02T15:04:05"

// StringToTime convert string to time
func StringToTime(value string) time.Time {
	var convertedTime, _ = time.Parse(layout, value)
	return convertedTime
}
