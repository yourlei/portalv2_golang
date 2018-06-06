package controller

import (
	"encoding/json"
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

// 查询用户与列表
func QueryUser(c *gin.Context) {
	var (
		code = 0
		queryJson common.UserQueryBody
	)
	// json string 转为 struct
	if err := json.Unmarshal([]byte(c.Query("query")), &queryJson); err != nil {
		common.RespondBadRequest(c)
		return
	}
	res, msg := service.QueryUserList(queryJson)
	if msg != nil {
		code = 1
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": msg,
		},
		"data": res,
		"total": len(res),
	})
}
// 更新用户状态
func UpdateUserStatus(c *gin.Context) {
	// post request body
	type JsonBody struct {
		Status int    `json:"status" binding:"required"`
		Remark string `json:"remark" binding:"required"`
	}
	var (
		body JsonBody
	)
	if err := c.BindJSON(&body); err != nil {
		common.RespondBadRequest(c)
    return
	}
	code, errMsg := service.UpdateUserStatus(c.Param("id"), body.Status, body.Remark)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": errMsg,
		},
	})
}
// test query
func GetList(c *gin.Context) {
	// var tem = &UserQueryBody{}
	// query := c.Query("query")

	// if query == "" {
	// 	common.RespondBadRequest(c)
	// 	return
	// }
	// if err := json.Unmarshal([]byte(query), &tem); err == nil {
	// 	fmt.Println(tem.Where)

	// 	if len(tem.Where.Email) > 0 {
	// 		fmt.Printf("hello===========")
	// 	}
	// }
	// if tem.Where.Email != "" {
	// 	fmt.Println(tem.Where.Email)
	// }

	// convert map 
	// if tem["where"] != nil {
	// 	fmt.Print("==========================")
	// 	body := tem["where"]
	// 	// fmt.Println(body)
	// 	if tem["where"]["email"] != nil {
	// 		fmt.Println(tem["where"])
	// 	}
	// }
	// fmt.Println(tem)
	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 0,
	// 	"data": "",
	// })
}