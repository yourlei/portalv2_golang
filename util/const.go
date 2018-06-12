package util

import (
	"time"
)
// Default Time Format
var TimeFormat = "2006-01-02 15:04:05"
// Default value of time type
var DefaultTime, _ = time.Parse(TimeFormat, "0001-01-01 00:00:00")