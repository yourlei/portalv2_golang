package inter

import (
	"net/http"
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