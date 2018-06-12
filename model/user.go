package model
import (
	"time"
)
/*************************/
/********用户结构体*********/
/*************************/
// 用户列表字段
type User struct {
	Id          int       `json:"id"`            //用户ID
	Name        string    `json:"name"`          //用户名
	Role        int       `json:"role"`          //角色ID
	Mobile      string    `json:"mobile"`        //电话
	Email       string    `json:"email"`         //电子邮箱
	Password    string    `json:"-"`             //密码
	Status      int       `json:"status"`        //状态
	CheckStatus int       `json:"check_status"`  //审核状态
	Remark      string    `json:"remark"`        //描述
	CheckRemark string    `json:"check_remark"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
// 登录返回的token
type TokenAndUser struct {
	Id          int       `json:"id"`            //用户ID
	RoleId      int       `json:"role_id"`       //角色id
	Name        string    `json:"name"`          //用户名
	Email       string    `json:"email"`         //电子邮箱
	Mobile      string    `json:"mobile"`        //电话
	Password    string    `json:"password"`      //密码
	Status      int       `json:"status"`        //状态
	CheckStatus int       `json:"check_status"`  //审核状态
	Token       string    `json:"token"`         //token
}
// 用户注册提交表单信息
type SignupForm struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Mobile    string `json:"mobile" binding:"required,len=11"`
	Password  string `json:"password" binding:"required,min=6,max=12"`
	RoleId    int    `json:"roleId" binding:"required,gt=0"`
}
// 登录时提交的请求体
type LoginForm struct {
	Email 	 string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=12"`
	Uuid     string `json:"uuid" binding:"required"`
	Code     string `json:"code" binding:"required"`
}
// 编辑用户
type EditUserForm struct {
	Name     string `json:"name,omitempty"`
	Mobile   string `json:"mobile,omitempty" binding:"len=11"`
	Password string `json:"password,omitempty" binding:"min=6,max=12"`
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