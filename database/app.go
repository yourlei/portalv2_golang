package database

import (
	"time"
	"strings"
	"portal/util"
	"portal/model"
	"github.com/satori/go.uuid"
)
// Create App
var createApp = `INSERT INTO portal_app(uuid, name, created_at, updated_at) VALUES(?, ?, ?, ?)`
// Select row
var selectApp = `SELECT id, name, uuid, created_at, updated_at FROM portal_app WHERE `
// Add app row
func CreateApp(name string) (int, error) {
	tx, err := ConnDB().Begin()

	if err != nil {
		return 1, err
	}
	// create uuid
	u1 := uuid.Must(uuid.NewV4())
	u2 := strings.Replace(u1.String(), "-", "", -1)
	
	_, err = tx.Exec(createApp, u2, name, time.Now().Format(util.TimeFormat), time.Now().Format(util.TimeFormat))
	if err != nil {
		return 1, err
	}
	tx.Commit()
	return 0, nil
}
// Search app list
func FindAllApp(where string, query ...interface{}) ([]interface{}, error) {
	var result = make([]interface{}, 0)
	rows, err := ConnDB().Query(selectApp + where, query...)

	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var ele = model.App{}
		if err := rows.Scan(
			&ele.Id,
			&ele.Name,
			&ele.Uuid,
			&ele.CreatedAt,
			&ele.UpdatedAt,
		); err != nil {
			return result, err
		} else {
			result = append(result, ele)
		}
	}
	return result, nil
}
// Edit app
func UpdateApp(id int, name string) (int, error) {
	Sql := `UPDATE portal_app SET name = ?, updated_at = ? WHERE id = ?`
	_, err := ConnDB().Exec(Sql, name, time.Now().Format(util.TimeFormat), id)
	if err != nil {
		return 1, err
	}
	return 0, nil
}