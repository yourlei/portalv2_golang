package model
import (
	"time"
)

const DeletedAt = "0000-01-01 00:00:00"

type QueryParams struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
// define time range
type DateRang struct {
	Gt time.Time  `json:"$gt,omitempty" binding:"required"`
	Lt time.Time  `json:"$lt,omitempty" binding:"required"`
}