package controller

import (
	"net/http"
	"fmt"

	"portal/service"
	"portal/common"

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

func Signup(c *gin.Context) {
	var signupInfo common.SignupForm
	var msg string

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
	code := service.Signup(signupInfo)

	switch code {
	case 100010:
		msg = "该邮箱或手机号已注册"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": msg,
		},
	})
}