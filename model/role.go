package model
import (
	"time"
)

//角色结构体
type Role struct {
	Id        int       `json:"id" binding:"min=1"`    //角色ID
	Name      string    `json:"name"`                  //角色名
	Remark    string    `json:"remark"`                //描述
	Status    int       `json:"status" binding:"min=1"`//状态
	CreatedAt time.Time `json:"created_at"`            //创建时间
	UpdatedAt time.Time `json:"updated_at"`            //更新时间
}