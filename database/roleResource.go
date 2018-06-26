package database

import (
	"portal/model"
)
// 查看角色的菜单资源
var roleMenuSql = "SELECT" +
										" portal_resource.id AS id," +
										" portal_router.name," +
										" portal_app.name  AS app" +
									" FROM" +
										" portal_resource" +
										" JOIN portal_app ON portal_resource.`app_id` = portal_app.`uuid`" + 
										" JOIN portal_router ON portal_resource.resource_id = portal_router.id" + 
									" WHERE" +
									" portal_resource.id IN ( SELECT resource_id FROM portal_role_res WHERE role_id = ? ) AND `portal_resource`.`type` = 1"
// 查看角色的接口资源
var roleInterSql =  "SELECT" +
											" portal_resource.id AS id," +
											" portal_interface.name," +
											" portal_app.name  AS app" +
										" FROM" +
											" portal_resource" +
											" JOIN portal_app on portal_resource.`app_id` = portal_app.`uuid`" + 
											" JOIN portal_interface ON portal_resource.resource_id = portal_interface.id" + 
										" WHERE" +
										" portal_resource.id IN ( SELECT resource_id FROM portal_role_res WHERE role_id = ? ) AND portal_resource.`type` = 2"
// 权限分配
func BindRoleRes(arg model.RolePrivilege) error {
	insertSql := "INSERT INTO portal_role_res(role_id, resource_id) VALUES(?, ?)"
	delSql    := "DELETE FROM portal_role_res WHERE role_id = ? AND resource_id = ?"
	sqlList   := []string{insertSql, delSql}

	tx, err := ConnDB().Begin()
	if err != nil {
		return err
	}
	for i := 0; i < len(sqlList); i++ {
		var d []int
		if i == 0 {
			d = arg.Enable
		} else {
			d = arg.Disable
		}
		// 每个menu | interface执行sql
		for _, v := range d {
			_, err := tx.Exec(sqlList[i], arg.Id, v)
			if err != nil {
				return err
			}
		}
	}
	return	tx.Commit()
}
// Get role permission
func GetRolePermmison(roleId int) (*model.ResBelongRole, error) {
	var (
		sqlList = []string{roleMenuSql, roleInterSql}
		menu  []model.ResourceInfo
		inter []model.ResourceInfo
	)
	for i, v := range sqlList {
		var rows, err = ConnDB().Query(v, roleId)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var ele = model.ResourceInfo{}
			if err := rows.Scan(
				&ele.Id,
				&ele.Name,
				&ele.App,
			);  err != nil {
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
	result := &model.ResBelongRole{Menu: menu, Inter: inter}
	return result, nil
}