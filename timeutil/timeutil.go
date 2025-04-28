package timeutil

import (
	"time"
)

// DayDiff return day diff correctly from two time
func DayDiff(start, end time.Time) int {
	startZero := start.Truncate(time.Hour * 24)
	endZero := end.Truncate(time.Hour * 24)
	return int(endZero.Sub(startZero).Hours() / 24)
}
