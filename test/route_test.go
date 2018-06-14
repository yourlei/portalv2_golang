package test

import (
	"testing"
	"portal/database"
)

func TestSetDeleteAt(t *testing.T) {
	database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
	_, err := database.SetDeletedAt(1, `portal_router`)

	if err != nil {
		t.Error(err)
	}
}