package controller

import (
	"net/http"
	"log"

	"portal/service"

	"github.com/gin-gonic/gin"
)
// 登录时提交的请求体结构
type loginForm struct {
	Email 	 string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Uuid     string `json:"uuid" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

// 注册时提交表单域
type signupForm struct {
	name      string `json:"name" binding:"required"`
	email     string `json:"email" binding:"required"`
	mobile    string `json:"mobile" binding:"required"`
	password  string `json:"password" binding:"required"`
	roleId    int    `json:"roleId" binding:"required"`
}
// 账户登录
func Signin(c *gin.Context) {
	var loginInfo loginForm
	
	if err := c.BindJSON(&loginInfo); err != nil {
  	// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code": 1,
		// 	"error": gin.H{
		// 		"msg": "参数错误",
		// 	},
		// })
		respondBadRequest(c)
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
	log.Print("begin....")
	service.Signin(loginInfo.Email, loginInfo.Password)
}

func Signup(c *gin.Context) {
	var signupInfo signupForm

	if err := c.BindJSON(&signupInfo); err != nil {
  	// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code": 1,
		// 	"error": gin.H{
		// 		"msg": "参数错误",
		// 	},
		// })
		respondBadRequest(c)
    return
	}

	log.Print(signupInfo.name)
}

func respondBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 1,
		"error": gin.H{
			"msg": "参数错误",
		},
	})
}