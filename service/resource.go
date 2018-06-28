// Resource module
package service

import (
	"portal/model"
	"portal/database"
)
// Return menu and resource info
func GetResource() (interface{}, error) {
	// Return Data format
	type Data struct {
		Menus []interface{} `json:"menus"`
		Interfaces []model.ResCollection `json:"interface"`
	}
	var (
		parentGroup []model.ResCollection
	  menuList []interface{}
	)
	MixResource, err := database.FindAllResource()
	if err != nil {
		return nil, err
	}
	menu := MixResource.Menu
	// filter parent menu
	for _, val := range menu {
		if val.ParentId == -1 {
			parentGroup = append(parentGroup, val)
		}
	}
	// push menu group by parent
	// 子菜单归到父级菜单下
	for _, val := range parentGroup {
		var group []model.ResCollection
		group = append(group, val)
		for _, ele := range menu {
			if ele.ParentId != -1 && ele.ParentId == val.DetailId {
				group = append(group, ele)
			}
		}
		menuList = append(menuList, group)
	}
	var res = Data{Menus: menuList, Interfaces: MixResource.Inter}
	// return result, nil
	return res, nil
}