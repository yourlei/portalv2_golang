package test

import (
	"fmt"
	"testing"
	"time"
	// "portal/common"
)

type DateRange struct {
	Gt time.Time `json:"$gt"`
	Lt time.Time `json:"$lt"`
}
type whereBody struct {
	CreatedAt DateRange `json:"created_at"`
	UpdatedAt DateRange `json:"updated_at"`
}
type query struct {
	where whereBody
}
func TestConst(t *testing.T)  {
	var where = &query{}
	defaultTime, _ := time.Parse("2006-01-02 15:04:05", "0001-01-01 00:00:00")
	if where.where.CreatedAt.Gt == defaultTime {
		fmt.Println("error")
	}
	// fmt.Println(where.where.CreatedAt.Gt)
}