package database
import (
	"portal/model"
)
var insertSql = "INSERT INTO portal_resource(`app_id`, `type`, `resource_id`) VALUES(?, ?, ?)"
var menuSql = "SELECT" +
								" r1.id AS DetailId," +
								" r1.name,"        +
								" r1.parent,"      +
							  " r3.app AS `group`,"  +
								" r2.type,"        +
								" r2.id AS RESID"  +
							" FROM" +
								" portal_router AS r1" +
								" JOIN portal_resource AS r2 ON r1.id = r2.resource_id" +
								" JOIN portal_app AS r3 ON r2.app_id = r3.uuid"
var interfaceSql = "SELECT" +
											" r1.id AS DetailId," + 
											" r1.name,"           +
											" -1 AS `parent`,"    +
											" r3.app AS `group`," +
										  " r2.type,"           +
											" r2.id AS RESID"     +
										" FROM" +
											" portal_interface AS r1" +
											" JOIN portal_resource AS r2 ON r1.id = r2.resource_id" +
											" JOIN portal_app AS r3 ON r2.app_id = r3.uuid"
                 
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
func FindAllResource() (*model.MixResource, error) {
	var (
		menu []model.ResCollection
		inter []model.ResCollection
		sqlList = []string{menuSql, interfaceSql}
	)
	// 查询Menu, Interface
	for i := 0; i < len(sqlList); i++ {
		rows, err := ConnDB().Query(sqlList[i])
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var ele = model.ResCollection{}
			if err := rows.Scan(
				&ele.DetailId,
				&ele.Name,
				&ele.ParentId,
				&ele.Group,
				&ele.ResType,
				&ele.RESId,
			); err != nil {
				return nil, err
			} else {
				switch i {
				case 0:
					menu = append(menu, ele)
				case 1:
					inter = append(inter, ele)
				}
			}
		}
	}
	result := &model.MixResource{Menu: menu, Inter: inter}
	return result, nil
}
// remove row
func DelResourceRow(resourceId int) (int, error) {
	_, err := ConnDB().Exec(`DELETE FROM portal_resource WHERE resource_id = ?`, resourceId)
	if err != nil {
		return 1, err
	}
	return 0, nil
}
