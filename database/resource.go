package database
import (
	"portal/model"
)
var insertSql = "INSERT INTO portal_resource(`app_id`, `type`, `resource_id`) VALUES(?, ?, ?)"
var resourceSql = ""
// Insert record
// menu,interface记录关联到资源管理表
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
// Search menu, interface, return menu, interface 
// data list
func FindAllResource() {
	var (
		// menus = make([]interface{}, 0)
	)
}