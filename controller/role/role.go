package role

import (
	"fmt"
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
		fmt.Println(err)
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
// Query role list
func QueryRoleList(c *gin.Context) {
	var (
		queryJson *model.RoleQueryBody
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