package database
import (
	"portal/model"
)
var insertSql = "INSERT INTO portal_resource(`app_id`, `type`, `resource_id`) VALUES(?, ?, ?)"
var menuSql = "SELECT" +
										" r1.id AS ROUID," +
										" r1.name," +
										" r1.parent," +
										" r2.type," +
										" r2.id AS RESID" +
								  " FROM" +
									  " portal_router AS r1" +
									  " JOIN portal_resource AS r2 ON r1.id = r2.resource_id"
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
func FindAllResource() ([]model.Menu2Res, error) {
	var (
		result = make([]model.Menu2Res, 0)
	)
	rows, err := ConnDB().Query(menuSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ele = model.Menu2Res{}
		if err := rows.Scan(
			&ele.RouId,
			&ele.Name,
			&ele.ParentId,
			&ele.ResType,
			&ele.RESId,
		); err != nil {
			return result, err
		} else {
			result = append(result, ele)
		}
	}
	return result, nil
}