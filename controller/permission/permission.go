package permission

import (
	"strconv"
	"net/http"
	"portal/util"
	"portal/model"
	"portal/service"
	"github.com/gin-gonic/gin"
)
// Grant permission to role
func Grant(c *gin.Context) {
	var (
		code int
		jsonBody model.RolePrivilege
	)
	err := c.BindJSON(&jsonBody)
	if err != nil {
		util.RespondBadRequest(c)
		return
	}
	err = service.Grant(jsonBody)
	if err != nil {
		code = 1
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": err,
		},
	})
}
// Get role permission
// 查看角色组已分配的权限
func GetRolePermisson(c *gin.Context) {
	var code = 0
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.RespondBadRequest(c)
		return
	}
	res, err := service.GetRolePermisson(id)
	if err != nil {
		code = 1
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": err,
		},
		"data": res,
	})
}
