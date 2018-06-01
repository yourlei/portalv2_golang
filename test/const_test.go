package test

import (
	"testing"
	"fmt"
	"portal/common"
)

func TestConst(t *testing.T)  {
	codeMsg := &common.CodeMsg{
		// code: 0,
		// "error": {
		// 	"msg": "slfjdl",
		// },
	}
	codeMsg.code = 0
	
	fmt.Println(codeMsg)
	// if id != 1 {
	// 	t.Error("查询错误", id, name)
	// }
}