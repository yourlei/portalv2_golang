package service

import (
	"portal/util"
	"portal/model"
	"portal/database"
)
// Create app
func CreateApp(name string) (int, interface{}) {
	code, err := database.FindByName(name, `portal_app`)
	if code == 0 {
		return 1, "该应用已存在"
	}
	if _, err = database.CreateApp(name); err != nil {
		return 1, err
	}
	return 0, nil
}
// Get App list
func GetAppList(query *model.GlobalQueryBody) ([]interface{}, error) {
	_sql, params := util.ParseQueryBody(query, false)
	res, err := database.FindAllApp(_sql, params...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
// Update app
func UpateApp(id int, name string) (int, error) {
	return database.UpdateApp(id, name)
}
