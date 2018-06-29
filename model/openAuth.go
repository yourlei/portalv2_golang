// Define opneAuth
// 外部应用有验证模型
package model

// 请求时提交的相关参数
type OpenAuth struct {
	Token  string  `json:"_"`
	Typeid int    `json:"typeid" binding:"required,min:1"`  // 验证方式
	Appid  string `json:"appid" binding:"required,min:1,max:32"` // 应用标识
}

