package database

import (
	"fmt"
)

func FindById(id string, table string) (bool, error) {
	var name string
	sql := "SELECT `name` FROM " + table + " WHERE id = ?"
	stmt, err := ConnDB().Prepare(sql)
	// IF error
	if err != nil {
		return false, err
	}
	err = stmt.QueryRow(id).Scan(&name)
	fmt.Println(name, "==========")
	if err != nil || name == "" {
		return false, err
	}
	return true, nil
}