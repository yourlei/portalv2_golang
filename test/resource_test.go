package test

import (
	"fmt"
	"testing"
	"portal/model"
	"portal/database"
)

func TestRou2Res(t *testing.T) {
	var (
		parentGroup []model.Menu2Res
		result []interface{}
	)
	database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
	list, _ := database.FindAllResource()
	// filter parent menu
	for _, val := range list {
		if val.ParentId == -1 {
			parentGroup = append(parentGroup, val)
		}
	}
	// push menu group by parent
	for _, val := range parentGroup {
		var group []model.Menu2Res
		group = append(group, val)
		for _, ele := range list {
			if ele.ParentId != -1 && ele.ParentId == val.RouId {
				group = append(group, ele)
			}
		}
		// result[index] = group
		result = append(result, group)
	}

	fmt.Println(result)
}