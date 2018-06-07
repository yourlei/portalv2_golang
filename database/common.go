package database

import(
	"database/sql"
)

func FindById(id string, table string) (int, error) {
	var (
		name string
		Sql = "SELECT `name` FROM " + table + " WHERE id = ?"
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