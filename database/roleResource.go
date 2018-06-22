package database

import (
	"portal/model"
)
// 权限分配
func BindRoleRes(arg model.RolePrivilege) error {
	insertSql := "INSERT INTO portal_role_res(role_id, resource_id) VALUES(?, ?)"
	delSql    := "DELETE FROM portal_role_res WHERE role_id = ? AND resource_id = ?"
	sqlList   := []string{insertSql, delSql}

	tx, err := ConnDB().Begin()
	if err != nil {
		return err
	}
	for i := 0; i < len(sqlList); i++ {
		var d []int
		if i == 0 {
			d = arg.Enable
		} else {
			d = arg.Disable
		}
		for _, v := range d {
			_, err := tx.Exec(sqlList[i], arg.Id, v)
			if err != nil {
				return err
			}
		}
	}
	return	tx.Commit()
}	