package test

import (
	"testing"
	"portal/model"
	"portal/database"
)

func TestSetDeleteAt(t *testing.T) {
	// database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
	// _, err := database.SetDeletedAt(1, `portal_router`)
	// res, err := database.ConnDB().Query(`SELECT name FROM portal_user WHERE name LIKE "%ad%"`)
	// result := make([]interface{}, 0)
	// type D struct {
	// 	name string
	// }
	// var dd D
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	for res.Next() {
	// 		if err := res.Scan(&dd.name); err != nil {
	// 			t.Error(err)
	// 		} else {
	// 			result = append(result, dd)
	// 		}
	// 	}
	// }
	// t.Log(result)
}
func TestUpdateRoute(t *testing.T) {
	database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
	code, msg := database.UpdateRouter(15, model.RouteUpdate{Name: "深圳行_世界之窗", Route: "world_window", Priority: 6, Schema: {"addr": "广州"}, Parent: 12})
	if code != 0 {
		t.Error(msg, "error")
	}
	t.Log(code, msg)
}