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
type GlobalQueryWhere struct {
	Name string `json:"name,omitempty"`
	App string `json:"app,omitempty"`
	CreatedAt DateRang `json:"created_at,omitempty"`
	UpdatedAt DateRang `json:"updated_at,omitempty"`
}
// global query struct
type GlobalQueryBody struct {
	QueryParams
	Where GlobalQueryWhere `json:"where"`
}