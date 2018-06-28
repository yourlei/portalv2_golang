// Role module
package service

import (
	"time"
	"portal/util"
	"portal/model"
	"portal/database"
)
// Create role
func CreateRole(name, remark string) (int, interface{})  {
	exsited, err := database.FindRoleByName(`name = ?`, name)
	if err == nil && exsited {
		return 20001, "该角色已存在"
	}
	err = database.CreateRole(name, remark)
	if err != nil {
		return 1, err
	}
	return 0, nil
}
// Update role info
func UpdateRole(id, name, remark string) (int, interface{}) {
	var (
		Sql string
		values []string
	)
	// check name exsited
	if name != "" {
		exsited, err := database.FindRoleByName(`name = ? AND id != ?`, name, id)
		if err == nil && exsited {
			return 20001, "该角色已存在"
		}
		Sql += `name = ?`
		values = append(values, name)
		// include remark
		if remark != "" {
			Sql += `, remark = ?`
			values = append(values, remark)
		}
	} else if remark != "" {
		Sql += `remark = ?`
		values = append(values, remark)
	}
	// name OR remark not nil
	if len(Sql) > 0 {
		Sql += `, updated_at = ? WHERE id = ?`
		values = append(values, time.Now().Format(util.TimeFormat), id)
		// string slice convert to interface slice
		params := make([]interface{}, len(values))
		for i, v := range values {
			params[i] = v
		}
		err := database.UpdateRole(Sql, params...)
		if err != nil {
			return 1, err
		}
	}
	return 0, nil
}
// Query role list
func GetRoleList(query *model.GlobalQueryBody) ([]interface{}, error) {
	_sql, params := util.ParseQueryBody(query, false)
	// Run sql
	res, err := database.FindAllRole(_sql, params...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
// Delete role
func DeleteRole(id int) (int, interface{}){
	code, _ := database.FindById(id, `portal_role`)
	if code != 0 {
		return 1, "角色不存在"
	}
	return database.DeleteRole(id)
}
// Get user on role group
func GetUserByRole(id int) ([]interface{}, error) {
 return	database.GetUserByRoleId(id)
}
// Migrate user to role group
func MigrateUser(roleId int, uids []int) error {
	return database.MigrateUser(roleId, uids)
}