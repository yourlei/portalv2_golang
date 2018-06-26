package service

import (
	"fmt"
	"strings"
	"encoding/json"

	"portal/util"
	"portal/model"
	"portal/database"
)

// Create Interface
func CreateInterface(Inter model.Interface) (int, interface{}) {
	// check name exsited
	code, _ := database.UniqueInterface(Inter.AppId, Inter.Name)
	if code != 0 {
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
// Show interface list
func GetInterfaceList(query *model.InterQueryBody) ([]interface{}, error) {
	var (
		where string = ""
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
	// Run sql
	res, err := database.FindAllInterface(where, params...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
// Edit interface
func UpdateInterface(id int, Inter model.InterfaceUpdate) (int, interface{}) {
	var (
		setSql []string
	)
	if Inter.Name != "" {
		code, _ := database.UniqueInterface(Inter.AppId, Inter.Name)
		if code != 0 {
			return 1, "接口已存在"
		}
		_sql := " `name` = '" + Inter.Name + "'"
		setSql = append(setSql, _sql)
	}
	if Inter.Schema != nil {
		schema, err := json.Marshal(Inter.Schema)
		if err != nil {
			return 1, err
		}
		_sql := " `schema` = '" + string(schema) + "'"
		setSql = append(setSql, _sql)
	}
	if Inter.Route != "" {
		_slq := " `route` = '" + Inter.Route + "'"
		setSql = append(setSql, _slq)
	}
	if Inter.Remark != "" {
		_slq := " `remark` = '" + Inter.Remark + "'"
		setSql = append(setSql, _slq)
	}
	var Sql = "UPDATE portal_interface SET " + strings.Join(setSql, ",") + " WHERE id = ?"
	fmt.Println(Sql)
	
	return database.EditInterface(Sql, id)
}
// Delete Interface
func DelInterface(id int) (int, interface{}) {
	code, msg := database.SetDeletedAt(id, `portal_interface`)
	if code != 0 {
		return code, msg
	}
	return database.DelResourceRow(id)
}
