package service

import (
	"portal/model"
	"portal/database"
)

// Create Interface
func CreateInterface(Inter model.Interface) (int, interface{}) {
	// check name exsited
	code, _ := database.FindByName(Inter.Name, `portal_interface`)
	if code == 0 {
		return 1, "接口已存在"
	}
	rowId, err := database.CreateInterface(Inter)
	if err != nil {
		return 1, err
	}
	// add record to resource table
	resource := model.Resource{ResType: 2, AppId: Inter.AppId, ResId: rowId}
	err = database.InsertRes(resource)
	if err != nil {
		return 1, err
	}
	return 0, nil
}