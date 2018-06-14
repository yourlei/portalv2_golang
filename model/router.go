package model

import (
	"time"
)
/*************************/
/********路由结构体*********/
/*************************/
type Route struct {
	Id        int       `json:"id"`                              //ID
	Name      string    `json:"name" binding:"required,min=1"`   //名称
	Route     string     `json:"item" binding:"required"`      //路由地址
	Type      int       `json:"action" binding:"required,min=1"`     //路由类型
	Parent    int       `json:"parent" binding:"required"`           //父级id
	Priority  int       `json:"priority" binding:"required"`    //权重
	Schema    interface{}    `json:"schema"`                         //参数配置
	Remark    string    `json:"remark"`                         //描述
	CreatedAt time.Time `json:"created_at"`                     //创建时间
	UpdatedAt time.Time `json:"updated_at"`                     //更新时间
}