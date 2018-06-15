package test

import (
	"testing"
	"portal/database"
)

func TestSetDeleteAt(t *testing.T) {
	database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
	// _, err := database.SetDeletedAt(1, `portal_router`)
	res, err := database.ConnDB().Query(`SELECT name FROM portal_user WHERE name LIKE "%ad%"`)
	result := make([]interface{}, 0)
	type D struct {
		name string
	}
	var dd D
	if err != nil {
		t.Error(err)
	} else {
		for res.Next() {
			if err := res.Scan(&dd.name); err != nil {
				t.Error(err)
			} else {
				result = append(result, dd)
			}
		}
	}
	t.Log(result)
}