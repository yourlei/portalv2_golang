// Interface module
// 提供与接口资源相关的功能函数
package service

import (
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
func GetInterfaceList(query *model.GlobalQueryBody) ([]interface{}, error) {
	_sql, params := util.ParseQueryBody(query, true)
	// Run sql
	res, err := database.FindAllInterface(_sql, params...)
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
	if len(setSql) > 0 {
		var Sql = "UPDATE portal_interface SET " + strings.Join(setSql, ",") + " WHERE id = ?"
		return database.EditInterface(Sql, id)
	}
	return 0, nil
}
// Delete Interface
func DelInterface(id int) (int, interface{}) {
	code, msg := database.SetDeletedAt(id, `portal_interface`)
	if code != 0 {
		return code, msg
	}
	return database.DelResourceRow(id)
}
