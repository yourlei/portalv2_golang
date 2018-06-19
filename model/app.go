package model

import (
	"time"
)
// 应用表,菜单可选择所归属的应用
type App struct {
	Id        int        //`json:"id"`
	Uuid      string     `json:"uuid,omitempty"`
	Name      string     `json:"name,omitempty"`
	CreatedAt time.Time  //`json:"created_at"`
	UpdatedAt time.Time  //`json:"updated_at"`
}
// App query 
type AppWhere struct {
	Name string `json:"name,omitempty"`
	CreatedAt DateRang `json:"created_at,omitempty"`
	UpdatedAt DateRang `json:"updated_at,omitempty"`
}
type AppQueryBody struct {
	QueryParams
	Where RouteWhere `json:"where"`
}