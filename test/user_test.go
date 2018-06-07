package test

import (
	// "fmt"
	"testing"
	"portal/service"
	"portal/database"
)

// func TestSignin(t *testing.T)  {
// 	id, name := service.Signin("test@qq.com", "123456")

// 	if id != 1 {
// 		t.Error("查询错误", id, name)
// 	}
// }

// func TestQueryUser(t *testing.T)  {
// 	res, err := service.QueryUserList()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Error(res)
// }

func TestChangePasswd(t *testing.T) {
	database.OpenDB("portal:D024Ad41d8cd98f00b204@tcp(192.168.80.243:3306)/portal?parseTime=true")
	// var id = "35"
	// s, err := database.FindById(id, "portal_users")
	// pw, _ := database.GetPasswd("35")
	code, err := service.ChangePasswd("37", "scut2017", "scut2019")
	if code != 0 {
		t.Error(err)
	}
}