package util

import (
	// "fmt"
	"time"
	"portal/model"
)
// Compare time range
func CompareDate(start, end time.Time) bool {
	 return start.Unix() > end.Unix()
}
// 解析查询参数, 输出对应的sql
func ParseQueryBody(query *model.GlobalQueryBody, where string) (string, []interface{}) {
	var (
		values []string
	)
	// include name
	if query.Where.Name != "" {
		where += ` AND name LIKE "%` + query.Where.Name + `%"`
	}
	// include created_at
	if query.Where.CreatedAt.Gt != DefaultTime {
		where += ` AND created_at BETWEEN ? AND ?`
		values = append(values, query.Where.CreatedAt.Gt.Format(TimeFormat), query.Where.CreatedAt.Lt.Format(TimeFormat))
	}
	// include updated_at
	if query.Where.UpdatedAt.Gt != DefaultTime {
		where += ` AND updated_at BETWEEN ? AND ?`
		values = append(values, query.Where.UpdatedAt.Gt.Format(TimeFormat), query.Where.UpdatedAt.Lt.Format(TimeFormat))
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

	return where, params
}