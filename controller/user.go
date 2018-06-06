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
	// request body
	type JsonBody struct {
		Status int    `json:"status" binding:"required"`
		Remark string `json:"remark" binding:"required"`
	}
	var body JsonBody
	if err := c.BindJSON(&body); err != nil {
		common.RespondBadRequest(c)
    return
	}
	code, errMsg := service.UpdateUserStatus(c.Param("id"), body.Status, body.Remark)
	r := &common.BaseResponse{
		Code: code,
	}
	r.Error.Msg = errMsg
	c.JSON(http.StatusOK, r)
}
// 审核用户
func ReviewUser(c *gin.Context) {
	// request body
	type JsonBody struct {
		Status int    `json:"check_status" binding:"required"`
		Remark string `json:"check_remark" binding:"required"`
	}
	var body JsonBody
	if err := c.BindJSON(&body); err != nil {
		common.RespondBadRequest(c)
    return
	}
	code, errMsg := service.ReviewUser(c.Param("id"), body.Status, body.Remark)
	r := &common.BaseResponse{
		Code: code,
	}
	r.Error.Msg = errMsg
	c.JSON(http.StatusOK, r)
}
// EditUser
func EditUser(c *gin.Context) {
	var body common.EditUserForm

	if err := c.BindJSON(&body); err != nil {
		common.RespondBadRequest(c)
    return
	}
	code, errMsg := service.EditUser(c.Param("id"), body)
	r := &common.BaseResponse{
		Code: code,
	}
	r.Error.Msg = errMsg
	c.JSON(http.StatusOK, r)
}
// test query
func Test(c *gin.Context) {
	fmt.Println("hi girl")

	// type responseBody struct{
	// 	Code int    `json:"code"`
	// 	Msg  error  `json:"msg"`
	// }
	// var str = `{"code": 0, "msg": "fail"}`
	// var d responseBody

	// var dat = &responseBody{Code: 0, Msg: errors.New("error down")}

	// if err := json.Unmarshal([]byte(str), &d); err != nil {
	// 	fmt.Println(err)
	// }
	// c.JSON(http.StatusOK, dat)
}