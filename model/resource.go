package model
// 资源管理表
// 连接不同类型的资源(如菜单,接口),角色的权限仅映射到该表的id
type Resource struct {
	Id      int 
	AppId   string //应用id
	ResType int   //资源类型
	ResId   int  //资源id
}
// 菜单,资源详情
type ResCollection struct {
	DetailId int     `json:"-"`//菜单id
	Name     string  `json:"name"`//菜单名称
	ResType  int     `json:"-"`//资源类型
	ParentId int     `json:"-"`//父级id
	RESId    int     `json:"id"`//关联资源id
	Group    string  `json:"group"`//所在模块
}
// Mix menu and interface resource
type MixResource struct {
	Menu  []ResCollection `json:"menus"`
	Inter []ResCollection `json:"interfaces"`
}