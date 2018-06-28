package model

import (
	"time"
)
type Interface struct {
	Id        int         `json:"id"` //ID         
	AppId     string      `json:"appid" binding:"required,max=32"` //所属应用                    
	Name      string      `json:"name" binding:"required,min=1"`   //名称
	// Group     string      `json:"group" binding:"required"`        //所属模块
	Route     string      `json:"route" binding:"required,min=1"`  //地址
	Schema    interface{} `json:"schema"`                          //参数配置
	Remark    string      `json:"remark"`                          //描述
	CreatedAt time.Time   `json:"created_at"`                      //创建时间
	UpdatedAt time.Time   `json:"updated_at"`                      //更新时间
}
// Interface Router query condition
// type InterWhere struct {
// 	Name      string   `json:"name,omitempty"`
// 	CreatedAt DateRang `json:"created_at,omitempty"`
// 	UpdatedAt DateRang `json:"updated_at,omitempty"`
// }
// // 接口列表查询体参数
// // Search menu list by condtion
// type InterQueryBody struct {
// 	QueryParams
// 	Where InterWhere `json:"where"`
// }
// 更新menu router
type InterfaceUpdate struct {
	Name     string      `json:"name"`
	Route    string      `json:"route"`
	AppId    string      `json:"appid" binding:"required,max=32"`
	Remark   string      `json:"remark"`
	Schema   interface{} `json:"schema"`
}