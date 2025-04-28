package timeutil

import (
	"time"
)

// DayDiff return day diff correctly from two time
func DayDiff(start, end time.Time) int {
    startZero := time.Date(
        start.Year(), start.Month(), start.Day(),
        0, 0, 0, 0,
        start.Location(),
    )
    endZero := time.Date(
        end.Year(), end.Month(), end.Day(),
        0, 0, 0, 0,
        end.Location(),
    )
	return int(endZero.Sub(startZero).Hours() / 24)
}
