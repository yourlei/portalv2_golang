package database
import (
	"portal/model"
)
var insertSql = "INSERT INTO portal_resource(`app_id`, `type`, `resource_id`) VALUES(?, ?, ?)"
// Insert record
func InsertRes(params model.Resource) error {
	tx, err := ConnDB().Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(insertSql, params.AppId, params.ResType, params.ResId)

	if err != nil {
		return err
	}
	return tx.Commit()
}