package test

import (
	"fmt"
	"testing"
	"portal/database"
)
// // Test method FindeRoleByName of database packe
// func TestFindRoleByName(t *testing.T) {
// 	database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
// 	code, err := database.FindRoleByName(`name = ?`, "研")

// 	if err != nil {
// 		t.Error(err)
// 	} else {
// 		t.Log("code: ", code)
// 	}
// }
// // Test CreateRole
// func TestCreateRole(t *testing.T) {
// 	database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
// 	code, err := database.CreateRole("测试", "测试用户组")

// 	if err != nil {
// 		t.Error(err)
// 	} else {
// 		t.Log("code: ", code)
// 	}
// }
// // Test Delete role
// func TestDeleteRole(t *testing.T) {
// 	database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
// 	code, err := database.DeleteRole(6)

// 	if err != nil {
// 		t.Error(err)
// 	} else {
// 		t.Log("code: ", code)
// 	}
// }
// // Update role 
// func TestUpdateRole(t *testing.T) {
// 	database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
// 	err := database.UpdateRole(6, "更新名称", "更新")

// 	if err != nil {
// 		t.Error(err)
// 	} else {
// 		t.Log("success")
// 	}
// }

func TestMigrateUser(t *testing.T) {
	database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
	var idGroup = make([]int, 0)
	idGroup =  append(idGroup, 1)
	
	fmt.Println(idGroup)
	err := database.MigrateUser(1, idGroup)

	if err != nil {
		t.Error(err)
	}
}