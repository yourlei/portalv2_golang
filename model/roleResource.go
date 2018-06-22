package model
// 角色,资源中间表
type RoleResource struct {
	Id         int
	RoleId     int
	ResourceId int
}
// type enableList struct {
// 	Enable  []int
// 	Disable []int
// }
//
// type RolePrivilege struct {
// Menu       enableList `json:"menus"`
// Interfaces enableList `json:"interfaces"`
// }
type RolePrivilege struct {
	Id int        `json:"id" bind:"required,min=1"`
	Enable  []int `json:"enable" binding:"required"`
	Disable []int `json:"disable" binding:"required"`
}