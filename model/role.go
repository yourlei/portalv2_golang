package model
import (
	"time"
)
/*************************/
/********角色结构体*********/
/*************************/
// Role Object
type Role struct {
	Id        int       `json:"id" binding:"min=1"`    //角色ID
	Name      string    `json:"name" binding:"required,min=1"` //角色名
	Remark    string    `json:"remark"`      //描述
	Status    int       `json:"-" binding:"min=1"`     //状态
	CreatedAt time.Time `json:"created_at"`            //创建时间
	UpdatedAt time.Time `json:"updated_at"`            //更新时间
}
// where 查询参数
// role where
// type RoleWhere struct {
// 	Name      string    `json:"name,omitempty"`
// 	CreatedAt DateRang  `json:"created_at,omitempty"`            //创建时间
// 	UpdatedAt DateRang  `json:"updated_at,omitempty"`            //更新时间
// }
// // role query body
// type RoleQueryBody struct {
// 	QueryParams
// 	Where RoleWhere `json:"where"`
// }