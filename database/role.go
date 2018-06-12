package database

import (
	"time"
	"fmt"
	"portal/util"
	"portal/model"
)

var createRole = `INSERT INTO portal_role(name, remark, created_at, updated_at) VALUES(?, ?, ?, ?)`
var selectRole = `SELECT id, name, remark, created_at, updated_at FROM portal_role WHERE `
// Create role
func CreateRole(name, remark string) (int64, error) {
	tx, err := ConnDB().Begin()
	if err != nil {
		return 1, err
	}

	res, err := tx.Exec(createRole, name, remark, time.Now().Format(util.TimeFormat), time.Now().Format(util.TimeFormat))
	if err != nil {
		return 1, err
	}
	id, _ := res.LastInsertId()
	tx.Commit()
	return id, nil
}
// Find Row by role name
func FindRoleByName(where string, query...interface{}) (bool, error) {
	type user struct {
		id int
		name string
	}
	var list = make([]interface{}, 0)
	Sql := `SELECT id, name FROM portal_role WHERE status = 1 AND `
	res, err := ConnDB().Query(Sql+where, query...)
	if err != nil {
		return false, err
	}
	for res.Next() {
		var ele = &user{}
		if err := res.Scan(
			&ele.id,
			&ele.name,
		); err != nil {
			return false, err
		} else {
			list = append(list, ele)
		}
	}
	
	return len(list) > 0, nil
}
// Update Role Info
func UpdateRole(id int, name, remark string) error {
	var Sql string
	if remark != "" {
		Sql = `UPDATE portal_role SET name =` + 
					`"` + name + `"` + `, remark = ` + 
					`"` + remark + `"` + `, updated_at = ? WHERE id = ?`
	} else {
		Sql = `UPDATE portal_role SET name = ` +
					`"` + name + `"` +
		      `, deleted_at =　? WHERE id = ?`
	}
	_, err := ConnDB().Exec(Sql, time.Now().Format(util.TimeFormat), id)
	if err != nil {
		return err
	}
	return nil
}
// Delete Role, set status = 2 
func DeleteRole(id int) (bool, error) {
	Sql := `UPDATE portal_role SET status = ?, deleted_at = ? WHERE id = ?`
	stmt, err := ConnDB().Prepare(Sql)
	if err != nil {
		return false, err
	}
	// exec sql
	_, err = stmt.Exec(2, time.Now().Format(util.TimeFormat), id)
	if err != nil {
		return false, err
	}
	return true, nil
}
// Find All User, Return Role List
func FindAllRole(where string, query ...interface{}) ([]interface{}, error) {
	var result = make([]interface{}, 0)
	rows, err := ConnDB().Query(selectRole + where, query...)

	if err != nil {
		fmt.Println(err, "one")
		return nil, err
	}
	defer rows.Close()
	// 遍历行, 追加到result slice
	for rows.Next() {
		var	data = &model.Role{}
		if err = rows.Scan(
			&data.Id,
			&data.Name,
			&data.Remark,
			&data.CreatedAt,
			&data.UpdatedAt,
		); err != nil {
			return result, err
		} else {
			result = append(result, data)
		}
	}
	return result, nil
}
func test() {
	fmt.Println("role module")
}