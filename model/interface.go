package model

import (
	"time"
)
type Interface struct {
	Id        int         `json:"id"`                              //ID
	Name      string      `json:"name" binding:"required,min=1"`   //名称
	Group     string      `json:"group" binding:"required"`        //所属模块
	Route     string      `json:"route" binding:"required,min=1"`  //地址
	Schema    interface{} `json:"schema"`                          //参数配置
	Remark    string      `json:"remark"`                          //描述
	CreatedAt time.Time   `json:"created_at"`                      //创建时间
	UpdatedAt time.Time   `json:"updated_at"`                      //更新时间
}