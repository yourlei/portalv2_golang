package model
// 资源管理表
// 连接不同类型的资源(如菜单,接口),角色的权限仅映射到该表的id
type Resource struct {
	Id      int 
	AppId   string //应用id
	resType int   //资源类型
	resId   int  //资源id
}