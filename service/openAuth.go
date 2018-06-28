// openAuth module
// 功能点:
// 1. 验证token
// 2. 验证用户访问数据接口权限
package service

import (
	"portal/util"
	"portal/model"
	"portal/config"
	"portal/database"
	"portal/middleware"
)
// Auth user permission and
// request logging
func Auth(req model.OpenAuth, log model.Log) (int, interface{}) {
	claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		return 1, "无效的token"
	}
	// parse token
	userId, err := util.Decrypt([]byte(config.AppConfig.AesKey), claims.UserId)
	roleId, err := util.Decrypt([]byte(config.AppConfig.AesKey), claims.RoleId) 
	if err != nil {
		return 1, "无效的token"
	}
	log.UserId = userId
	// loggin
	database.CreateLog(log)
	// return if typeid = 2
	if req.TypeId == 2 {
		return 0, nil
	}
	return database.CheckPermission(roleId, req.AppId)
}