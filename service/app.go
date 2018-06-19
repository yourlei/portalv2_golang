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
func GetAppList(query *model.AppQueryBody) ([]interface{}, error) {
	var (
		where = `deleted_at = "0000-01-01 00:00:00"`
		values []string
	)
	// include name
	if query.Where.Name != "" {
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

	res, err := database.FindAllApp(where, params...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
