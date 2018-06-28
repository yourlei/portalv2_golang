package service

import (
	"portal/util"
	"portal/model"
	"portal/database"
)
// Create app
func CreateApp(name string) (int, interface{}, string) {
	code, _ := database.UniqueAppName(name)
	if code == 0 {
		return 1, "该应用已存在", ""
	}
	_, uuid, err := database.CreateApp(name)
	if err != nil {
		return 1, err, ""
	}
	return 0, nil, uuid
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
