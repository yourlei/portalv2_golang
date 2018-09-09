package role

import (
	// "fmt"
	"strconv"
	"net/http"
	"encoding/json"

	"portal/util"
	"portal/model"
	"portal/service"

	"github.com/gin-gonic/gin"
)
// Create role
func CreateRole(c *gin.Context) {
	type role struct {
		Name   string `json:"name,omitempty" binding:"required"`
		Remark string `json:"remark"`
	}
	var jsonBody role
	err := c.BindJSON(&jsonBody)
	if err != nil {
		util.RespondBadRequest(c)
		return
	}
	code, msg := service.CreateRole(jsonBody.Name, jsonBody.Remark)
	r := &util.BaseResponse{
		Code: code,
	}
	r.Error.Msg = msg
	c.JSON(http.StatusOK, r)
}
// Update role
func UpdateRole(c *gin.Context) {
	type role struct {
		Name   string `json:"name,omitempty"`
		Remark string `json:"remark,omitempty"`
	}
	var jsonBody role
	err := c.BindJSON(&jsonBody)
	if err != nil {
		util.RespondBadRequest(c)
		return
	}
	code, msg := service.UpdateRole(c.Param("id"), jsonBody.Name, jsonBody.Remark)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": msg,
		},
	})
}
// Query role list
func QueryRoleList(c *gin.Context) {
	var (
		queryJson *model.GlobalQueryBody
		code int = 0
	)
	// json string 转为 struct
	if err := json.Unmarshal([]byte(c.Query("query")), &queryJson); err != nil {
		util.RespondBadRequest(c)
		return
	}
	// check time range
	where := queryJson.Where
	if util.CompareDate(where.CreatedAt.Gt, where.CreatedAt.Lt) ||
	   util.CompareDate(where.UpdatedAt.Gt, where.UpdatedAt.Lt) {
		util.RespondBadRequest(c)
		return	
	}
	res, msg := service.GetRoleList(queryJson)
	if msg != nil {
		code = 1
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": msg,
		},
		"data": res,
		"total": len(res),
	})
}
// Delete Role
func DeleteRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.RespondBadRequest(c)
		return
	}
	code, msg := service.DeleteRole(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": msg,
		},
	})
}
// Get user on role group
func GetUserByRole(c *gin.Context) {
	var code int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.RespondBadRequest(c)
		return
	}
	result, msg := service.GetUserByRole(id)
	if msg != nil {
		code = 1
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": msg,
		},
		"data": result, 
	})
}
// Migrate user 
func MigrateUser(c *gin.Context) {
	type Body struct {
		RoleId int   `json:"roleId" binding:"required"`
		UserId []int `json:"userId" binding:"required,min=1"`
	}
	var (
		jsonBody Body
		code int
	)
	err := c.BindJSON(&jsonBody)
	if err != nil {
		util.RespondBadRequest(c)
		return
	}
	err = service.MigrateUser(jsonBody.RoleId, jsonBody.UserId)
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