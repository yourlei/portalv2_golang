package model

//用户结构体
type User struct {
	ID          string //用户ID
	Name        string //用户名
	RoleID      string //角色ID
	Mobile      string //电话
	Email       string //电子邮箱
	Password    string //密码
	Status      string //状态
	CheckStatus string //审核状态
}