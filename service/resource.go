package service

import (
	"portal/model"
	"portal/database"
)
// Return menu and resource info
func GetResource() ([]interface{}, error) {
	var (
		parentGroup []model.Menu2Res
		result []interface{}
	)
	list, err := database.FindAllResource()
	if err != nil {
		return nil, err
	}
	// filter parent menu
	for _, val := range list {
		if val.ParentId == -1 {
			parentGroup = append(parentGroup, val)
		}
	}
	// push menu group by parent
	// 子菜单归到父级菜单下
	for _, val := range parentGroup {
		var group []model.Menu2Res
		group = append(group, val)
		for _, ele := range list {
			if ele.ParentId != -1 && ele.ParentId == val.RouId {
				group = append(group, ele)
			}
		}
		result = append(result, group)
	}
	return result, nil
}