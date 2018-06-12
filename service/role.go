package service

import (
	"fmt"
	// "fmt"
	"portal/util"
	"portal/model"
	"portal/database"
)
// Find User List
func GetRoleList(query *model.RoleQueryBody) ([]interface{}, error) {
	fmt.Println(query)
	var (
		where = `status = 1`
		values []string
	)
	// include name 
	if query.Where.Name != "" {
		where += ` AND name = ?`
		values = append(values, query.Where.Name)
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
	res, err := database.FindAllRole(where, params...)
	if err != nil {
		return nil, err
	}
	return res, nil
}