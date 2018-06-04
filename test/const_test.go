package test

import (
	"testing"
	"time"
	// "portal/common"
)

func TestConst(t *testing.T)  {
	// codeMsg := &common.CodeMsg{
		// code: 0,
		// "error": {
		// 	"msg": "slfjdl",
		// },
	// }
	// codeMsg.code = 0
	mytime := time.Now().Format("2006-01-02 15:04:05")
	t.Error(mytime)
	// if id != 1 {
	// 	t.Error("查询错误", id, name)
	// }
}