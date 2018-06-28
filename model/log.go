package model

import (
	// "time"
)
// Log model
type Log struct {
	Url      string    `json:"url"`           // 访问路径
	Method   string    `json:"method"`        // http方法
	Proto    string    `json:"proto"`         // 网络传输协议
	Agent    string    `json:"agent"`         // 用户使用的客户端
	Host     string    `json:"host"`          // 访问主机ip
	UserId   string    `json:"userid"`        // 用户id
	CreateAt string    `json:"created_at"`    // 访问时间
}