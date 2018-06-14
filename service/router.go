package service

import (
	"portal/model"
	"portal/database"
)

// Create router
func CreateRouter(r model.Route) (int, interface{}) {
	// check router uniqueness
	code, _ := database.UniqueRouter(r)
	if code == 0 {
		return 30001, "名称或地址已占用"
	}
	_, err := database.CreateRouter(r)
	if err != nil {
		return 1, err
	}
	return 0, nil
}
// Delete Route By id
func DeleteRoute(id int) (int, interface{}) {
	return database.SetDeletedAt(id, `portal_router`)
}