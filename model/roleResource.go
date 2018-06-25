package model
// 角色,资源中间表
type RoleResource struct {
	Id         int
	RoleId     int
	ResourceId int
}
// Assign permission to role
// Id: roleId
// Enable: grant permission
// Disable: revoke permission
type RolePrivilege struct {
	Id int        `json:"id" bind:"required,min=1"`
	Enable  []int `json:"enable" binding:"required"`
	Disable []int `json:"disable" binding:"required"`
}
type ResourceInfo struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	App  string   `json:"app"`
	// AppId string  `json:"appid"`
}
// Get belong to role's permission
// Result include route and interface
type ResBelongRole struct {
	Menu    []ResourceInfo `json:"menus"`
	Inter   []ResourceInfo `json:"interfaces"`
}