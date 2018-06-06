// 公用常量
package common

import (
	// "time"
)
/* *******************
 * 常量
 * ******************/
// deleted_at 默认值
const DeletedAt = "0000-01-01 00:00:00"

/* *******************
 * 结构体类
 * ******************/ 
// 用户注册提交表单信息
type SignupForm struct {
	Name      string `json:"name" binding:"required"`
	Email     string `valid:"email~请输入正确的邮箱" json:"email" binding:"required"`
	Mobile    string `valid:"numeric~请输入正确的手机号码" json:"mobile" binding:"required"`
	Password  string `json:"password" binding:"required"`
	RoleId    int    `json:"roleId" binding:"required"`
}
// 登录时提交的请求体
type LoginForm struct {
	Email 	 string `valid:"email~请输入正确的邮箱" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Uuid     string `json:"uuid" binding:"required"`
	Code     string `json:"code" binding:"required"`
}
/* *******************
 *    query body
 * ******************/
type QueryParams struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
// 用户列表查询条件
type UserWhere struct {
	Email       string    `json:"email,omitempty"`
	Mobile      string    `json:"mobile,omitempty"`
	Group       string    `json:"group,omitempty"`
	Status  	  string    `json:"status,omitempty"`
	CheckStatus string    `json:"check_status,omitempty"`
	CreatedAt   string    `json:"created_at,omitempty"`
}

type UserQueryBody struct {
	QueryParams
	Where UserWhere `json:"where"`
}