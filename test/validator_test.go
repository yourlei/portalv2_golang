package test

import (
	"fmt"
	"testing"
	"portal/common"
)
type User struct {
    Id    int    `validate:"number,min=1,max=1000"`
    Name  string `validate:"string,min=2,max=10"`
    Bio   string `validate:"string"`
    Email string `validate:"email"`
}

func TestIsMail(t *testing.T)  {
// 	result, msg := common.IsEmail("12Hello@qq.com")

// 	if !result {
// 		t.Error("验证失败", msg)
// 	}
		user := User {
			Id:    0,
			Name:  "superlongstring",
			Bio:   "",
			Email: "foobar",
    }

    fmt.Println("Errors:")
    for i, err := range common.ValidateStruct(user) {
        fmt.Printf("\t%d. %s\n", i+1, err.Error())
    }
}
