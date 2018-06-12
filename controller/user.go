package controller

import (
	"strconv"
	"encoding/json"
	"net/http"
	"fmt"

	"portal/util"
	"portal/model"
	"portal/config"
	"portal/service"

	"github.com/gin-gonic/gin"
)
// 账户登录
func Signin(c *gin.Context) {
	var loginInfo model.LoginForm
	
	if err := c.BindJSON(&loginInfo); err != nil {
		util.RespondBadRequest(c)
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
	code, data := service.Signin(loginInfo.Email, loginInfo.Password)
	msg, ok := data.(string)
	if !ok {
		// set cookie
		d, _ := data.(model.TokenAndUser)
		cookie := &http.Cookie{
			Name:     "token",
			Value:    d.Token,
			Path:     "/",
			HttpOnly: false,
			MaxAge:   config.AppConfig.TokenMaxAge,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"error": gin.H{
				"msg": "",
			},
			"data": data,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"error": gin.H{
				"msg": msg,
			},
		})
	}
}
// 用户注册
func Signup(c *gin.Context) {
	var signupInfo model.SignupForm
	// var msg string
	if err := c.BindJSON(&signupInfo); err != nil {
		util.RespondBadRequest(c)
    return
	}
	// code
	code, msg := service.Signup(signupInfo)
	r := &util.BaseResponse{
		Code: code,
	}
	r.Error.Msg = msg
	c.JSON(http.StatusOK, r)
}
// 查询用户与列表
func QueryUser(c *gin.Context) {
	var (
		code = 0
		queryJson model.UserQueryBody
	)
	// json string 转为 struct
	if err := json.Unmarshal([]byte(c.Query("query")), &queryJson); err != nil {
		util.RespondBadRequest(c)
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
	// id string convert int
	id, errStr := strconv.Atoi(c.Param("id"))
	err := c.BindJSON(&body)
	if errStr != nil || err != nil {
		util.RespondBadRequest(c)
    return
	}
	code, errMsg := service.UpdateUserStatus(id, body.Status, body.Remark)
	r := &util.BaseResponse{
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
	// id string convert int
	id, errStr := strconv.Atoi(c.Param("id"))
	err := c.BindJSON(&body)
	if errStr != nil || err != nil {
		util.RespondBadRequest(c)
    return
	}
	code, errMsg := service.ReviewUser(id, body.Status, body.Remark)
	r := &util.BaseResponse{
		Code: code,
	}
	r.Error.Msg = errMsg
	c.JSON(http.StatusOK, r)
}
// EditUser
func EditUser(c *gin.Context) {
	var body model.EditUserForm
	// id string convert int
	id, errStr := strconv.Atoi(c.Param("id"))
	err := c.BindJSON(&body)
	if errStr != nil || err != nil {
		util.RespondBadRequest(c)
    return
	}
	code, errMsg := service.EditUser(id, body)
	r := &util.BaseResponse{
		Code: code,
	}
	r.Error.Msg = errMsg
	c.JSON(http.StatusOK, r)
}
// Change password
func ChangePasswd(c *gin.Context) {
	type Password struct {
		OldPasswd string `json:"passwd"`
		NewPasswd string `json:"new_passwd"`
	}
	var body Password
	id, errStr := strconv.Atoi(c.Param("id"))
	err := c.BindJSON(&body)
	if errStr != nil || err != nil {
		util.RespondBadRequest(c)
    return
	}
	code, errMsg := service.ChangePasswd(id, body.OldPasswd, body.NewPasswd)
	r := &util.BaseResponse{
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