// Router module
// 菜单资源
package service

import (
	"portal/util"
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
	// check appid if parent not equal -1
	if r.Parent != -1 {
		equal, _ := database.EqualAppid(r.Parent, r.AppId)
		if !equal {
			return 30002, "所属应用与父菜单不一致"
		}
	}
	rowId, err := database.CreateRouter(r)
	if err != nil {
		return 1, err
	}
	resource := model.Resource{AppId: r.AppId, ResType: 1, ResId: rowId}
	// 关联resource表
	err = database.InsertRes(resource)
	if err != nil {
		return 1, err
	}
	return 0, nil
}
// Update menu
func UpdateRouter(id int, r model.RouteUpdate) (int, interface{}) {
	return database.UpdateRouter(id, r)
}
// Query menu router list
func GetRouterList(query *model.GlobalQueryBody) ([]interface{}, error) {
	_sql, params := util.ParseQueryBody(query, true)
	// Run sql
	res, err := database.FindAllRouter(_sql, params...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
// Delete Route By id
func DeleteRouter(id int) (int, interface{}) {
	return database.SetDeletedAt(id, `portal_router`)
}
// Get Parent route
func GetParentRoute() ([]interface{}, error) {
	return database.FindParentRouter()
}