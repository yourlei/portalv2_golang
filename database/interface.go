package database

import (
	"time"
	// "errors"
	"encoding/json"
	"portal/util"
	"portal/model"
)
var createInterface = "INSERT INTO `portal_interface`(`name`, `route`, `schema`, `remark`, `created_at`, `updated_at`) VALUES(?,?,?,?,?,?)"
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