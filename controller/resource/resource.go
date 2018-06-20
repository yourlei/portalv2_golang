package resource

import (
	"net/http"
	// "portal/model"
	"portal/service"
	"github.com/gin-gonic/gin"
)
func GetResource(c *gin.Context) {
	data, msg := service.GetResource()

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"error": gin.H{
			"msg": msg,
		},
		"data": data,
	})
}