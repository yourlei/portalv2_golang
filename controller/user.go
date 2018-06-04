package controller

import (
	"encoding/json"
	"net/http"
	"fmt"

	"portal/service"
	"portal/common"
	"portal/model"

	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
)
// 账户登录
func Signin(c *gin.Context) {
	var loginInfo common.LoginForm
	
	if err := c.BindJSON(&loginInfo); err != nil {
		common.RespondBadRequest(c)
    return
	}
	// 检查验证码
	if errCode, errMsg := service.VerifyCaptcha(loginInfo.Uuid, loginInfo.Code); errCode != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": errCode, 
			"error": gin.H{
				"msg": errMsg,
			},
		})
		return
	}
	service.Signin(loginInfo.Email, loginInfo.Password)
}
// 用户注册
func Signup(c *gin.Context) {
	var signupInfo common.SignupForm
	// var msg string

	if err := c.BindJSON(&signupInfo); err != nil {
		common.RespondBadRequest(c)
    return
	}
	// 验证注册信息
	if ok, err := govalidator.ValidateStruct(signupInfo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": err,
		})
	} else {
		fmt.Printf("OK: %v\n", ok)
	}
	// code
	code, msg := service.Signup(signupInfo)
	fmt.Println(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": error.Error(msg),
		},
	})
}

func QueryUser(c *gin.Context) {
	var code = 0
	res, msg := service.QueryUserList()
	fmt.Println(res, msg)
	var result = &model.UserList{Data: res}
	bts, _ := json.Marshal(result)
	if msg != nil {
		code = 1
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": error.Error(msg),
		},
		"data": bts,
	})
}