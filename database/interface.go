package database

import (
	"time"
	"database/sql"
	"encoding/json"
	"portal/util"
	"portal/model"
)
var createInterface = "INSERT INTO `portal_interface`(`name`, `route`, `schema`, `remark`, `created_at`, `updated_at`) VALUES(?,?,?,?,?,?)"
var selectInterface = "SELECT" +
												" a.id," +
												" a.`name`," +
												" a.route,"  +
												" a.remark," +
												" c.`name` AS app," +
												" a.created_at," +
												" a.updated_at" +
											" FROM" +
												" portal_interface AS a" +
												" JOIN portal_resource AS b ON a.id = b.resource_id" +
												" JOIN portal_app AS c ON b.app_id = c.uuid " +
											" WHERE "
// Create Router
func CreateInterface(Inter model.Interface) (int, error) {
	var v []byte
	tx, err := ConnDB().Begin()
	if err != nil {
		return 1, err
	}
	// interface to json string
	if v, err = json.Marshal(Inter.Schema); err != nil {
		return 1, err
	}
	res, err := tx.Exec(createInterface, Inter.Name, Inter.Route, string(v), Inter.Remark, time.Now().Format(util.TimeFormat), time.Now().Format(util.TimeFormat))
	if err != nil {
		return 1, err
	}
	err = tx.Commit()
	if err != nil {
		return 1, err
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}
// Find All interface
func FindAllInterface(where string, query...interface{}) ([]interface{}, error) {
	var (
		resutl []interface{}
	)
	rows, err := ConnDB().Query(selectInterface + where, query...)
	if err != nil {
		return nil , err
	}
	defer rows.Close()
	for rows.Next() {
		var ele = model.Interface{}
		if err := rows.Scan(
			&ele.Id,
			&ele.Name,
			&ele.Route,
			&ele.Remark,
			&ele.AppId,
			&ele.CreatedAt,
			&ele.UpdatedAt,
		); err != nil {
			return nil, err
		}
		resutl = append(resutl, ele)
	}
	return resutl, nil
}
// Edit interface row
func EditInterface(sql string, args...interface{}) (int, error) {
	_, err := ConnDB().Exec(sql, args...)
	if err != nil {
		return 1, err
	}
	return 0, nil
}
// check Unique
func UniqueInterface(query...interface{}) (int, error) {
	// 检查app组下接口名是否已占用
	var Sql = `SELECT COUNT(1) AS count FROM portal_interface AS a WHERE` +
						` a.id IN (SELECT resource_id FROM portal_resource WHERE app_id = ?)` +
						` AND a.name = ?`
  var count int
	err := ConnDB().QueryRow(Sql, query...).Scan(&count)
	// not found
	if err == sql.ErrNoRows {
		return 0, nil
	}
	// error
	if err != nil || count != 0 {
		return 1, err
	}
	return 0, nil
}