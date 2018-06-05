package model
import (
	// "time"
)
//用户结构体
// type User struct {
// 	ID          string //用户ID
// 	Name        string //用户名
// 	RoleID      string //角色ID
// 	Mobile      string //电话
// 	Email       string //电子邮箱
// 	Password    string //密码
// 	Status      string //状态
// 	CheckStatus string //审核状态
// }

type User struct {
	Id          string  `json:"id"`   //用户ID
	Name        string  `json:"name"` //用户名
	Role        int     `json:"role"`                //角色ID
	Mobile      string  `json:"mobile"`//电话
	Email       string  `json:"email"`//电子邮箱
	// Password    string //密码
	Status      int     `json:"status"`//状态
	CheckStatus int     `json:"check_status"`//审核状态
	// CreatedAt   time.Time
	// UpdatedAt   time.Time 
}

type UserList struct {
	Data []*User `json:"data"`
}