package database

import (
	// "database/sql"
	"errors"
)
func CheckPermission(roleid, appid string) (int, error) {
	var count int
	Sql := "SELECT" +
						" count(1)" +
					" FROM" +
						" portal_resource a"  +
					" WHERE" +
						" a.id IN (SELECT resource_id FROM portal_role_res WHERE role_id = ?)"  +
						" AND a.app_id = ?"
	err := ConnDB().QueryRow(Sql, roleid, appid).Scan(&count)
	// error
	if err != nil {
		return 1, err
	}
	// result
	if count > 0 {
		return 0, nil
	} else {
		return 1, errors.New("没有访问权限")
	}
}