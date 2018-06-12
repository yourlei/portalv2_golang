package util

import (
	"time"
)
// Compare time range
func CompareDate(start, end time.Time) bool {
	 return start.Unix() > end.Unix()
}