package inter

import (
	"strconv"
	"net/http"
	"encoding/json"

	"portal/util"
	"portal/model"
	"portal/service"

	"github.com/gin-gonic/gin"
)
// Create Interface
func CreateInterface(c *gin.Context) {
	var jsonBody model.Interface

	err := c.BindJSON(&jsonBody)
	if err != nil {
		util.RespondBadRequest(c)
		return
	}
	code, msg := service.CreateInterface(jsonBody)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": msg,
		},
	})
}
// Get interface list
func GetInterfaceList(c *gin.Context) {
	var (
		queryJson *model.InterQueryBody
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
	res, msg := service.GetInterfaceList(queryJson)
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
// Edit interface
func UpdateInterface(c *gin.Context) {
	var jsonBody model.InterfaceUpdate
	if err := c.BindJSON(&jsonBody); err != nil {
		util.RespondBadRequest(c)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.RespondBadRequest(c)
		return
	}
	code, msg := service.UpdateInterface(id, jsonBody)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": msg,
		},
	})
}
// Delete Interface
func DeleteInterface(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.RespondBadRequest(c)
		return
	}
	code, msg := service.DelInterface(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": msg,
		},
	})
}