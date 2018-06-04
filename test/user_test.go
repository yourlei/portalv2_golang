package test

import (
	"testing"
	"portal/service"
)

// func TestSignin(t *testing.T)  {
// 	id, name := service.Signin("test@qq.com", "123456")

// 	if id != 1 {
// 		t.Error("查询错误", id, name)
// 	}
// }

func TestQueryUser(t *testing.T)  {
	res, err := service.QueryUserList()

	if err != nil {
		t.Error(err)
	}

	t.Error(res)
}