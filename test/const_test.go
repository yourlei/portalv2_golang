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
// func TestConst(t *testing.T)  {
// 	var where = &query{}
// 	defaultTime, _ := time.Parse("2006-01-02 15:04:05", "0001-01-01 00:00:00")
// 	if where.where.CreatedAt.Gt == defaultTime {
// 		fmt.Println("error")
// 	}
// 	// fmt.Println(where.where.CreatedAt.Gt)
// }

// func TestParams(t *testing.T) {
// 	var obj []interface{}
// 	obj = append(obj, 1, 2, 3)

// 	var i64 int64 = 3222222222222222222	

// 	var i int = int(i64)
// 	fmt.Println(i64, i)
// }
func TestGetArgs(t *testing.T) {
	
}