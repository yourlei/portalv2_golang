package database

import(
	"database/sql"
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