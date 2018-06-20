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
// Query menu router list
func GetRouterList(query *model.RouteQueryBody) ([]interface{}, error) {
	var (
		where string = `deleted_at = "0000-01-01 00:00:00"`
		values []string
	)
	// include name 
	if query.Where.Name != "" {
		// where += ` AND name = ?`
		// values = append(values, query.Where.Name)
		where += ` AND name LIKE "%` + query.Where.Name + `%"`
	}
	// include created_at
	if query.Where.CreatedAt.Gt != util.DefaultTime {
		where += ` AND created_at BETWEEN ? AND ?`
		values = append(values, query.Where.CreatedAt.Gt.Format(util.TimeFormat), query.Where.CreatedAt.Lt.Format(util.TimeFormat))
	}
	// include updated_at
	if query.Where.UpdatedAt.Gt != util.DefaultTime {
		where += ` AND updated_at BETWEEN ? AND ?`
		values = append(values, query.Where.UpdatedAt.Gt.Format(util.TimeFormat), query.Where.UpdatedAt.Lt.Format(util.TimeFormat))
	}
	if query.Limit == 0 {
		query.Limit = 10
	}
	// Select offset and limit
	where += " LIMIT ?, ?"
	// slice不能直接传递给interface slice
	params := make([]interface{}, len(values)+2)
	for i, v := range values {
		params[i] = v
	}
	// 加入分页
	params[len(values)] = query.Offset
	params[len(values) + 1] = query.Limit
	// Run sql
	res, err := database.FindAllRouter(where, params...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
// Delete Route By id
func DeleteRoute(id int) (int, interface{}) {
	return database.SetDeletedAt(id, `portal_router`)
}