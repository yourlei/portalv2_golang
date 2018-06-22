package grant

import (
	"net/http"
	"portal/util"
	"portal/model"
	"portal/service"
	"github.com/gin-gonic/gin"
)
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