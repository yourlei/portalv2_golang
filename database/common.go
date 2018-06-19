package database

import(
	"time"
	"errors"
	"database/sql"
	"portal/util"
)
// Find a row by id
func FindById(id int, table string) (int, error) {
	var (
		name string
		Sql = "SELECT `name` FROM " + table + " WHERE id = ? AND deleted_at = '0000-01-01 00:00:00'"
	)
	err := ConnDB().QueryRow(Sql, id).Scan(&name)
	// not found
	if err == sql.ErrNoRows {
		return -1, nil
	}
	// error
	if err != nil {
		return 1, err
	}
	return 0, nil
}
// Set Deleted_at column in table
func SetDeletedAt(id int, table string) (int, interface{}) {
	var Sql = `UPDATE ` + table + ` SET deleted_at = ? WHERE id = ?`
	stmt, err := ConnDB().Prepare(Sql)
	if err != nil {
		return 1, err
	}
	res, err := stmt.Exec(time.Now().Format(util.TimeFormat), id)
	if err != nil {
		return 1, err
	}
	if effectId, _ := res.RowsAffected(); effectId == 0 {
		return 1, errors.New("id 不存在")
	}
	return 0, nil
}
// Find row by name
func FindByName(name, table string) (int, error) {
	var _name string
	Sql := `SELECT name FROM ` + table + ` WHERE name = ?`
	err := ConnDB().QueryRow(Sql, name).Scan(&_name)
	// not found
	if err == sql.ErrNoRows {
		return -1, nil
	}
	// error
	if err != nil {
		return 1, err
	}
	return 0, nil
}
