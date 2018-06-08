package model
import (
	"time"
)
// 用户结构体
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