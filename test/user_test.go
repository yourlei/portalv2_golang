package test

import (
	"testing"
	"portal/service"
)

func TestSignin(t *testing.T)  {
	id, name := service.Signin("test@qq.com", "123456")

	if id != 1 {
		t.Error("查询错误", id, name)
	}
}