package service

import (
	"fmt"
	"portal/model"
	"portal/database"
)
// Return menu and resource info
func GetResource() (interface{}, error) {
	var (
		parentGroup []model.ResCollection
		result []interface{}
	)
	MixResource, err := database.FindAllResource()
	if err != nil {
		fmt.Println(err)
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
		result = append(result, group)
	}
	type Data struct {
		Menus []interface{} `json:"menus"`
		Interfaces []model.ResCollection `json:"interface"`
	}
	// fmt.Println(result, )
	var res = Data{Menus: result, Interfaces: MixResource.Inter}
	// return result, nil
	return res, nil
}