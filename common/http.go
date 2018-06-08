package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
// 响应http 状态码400
func RespondBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 1,
		"error": gin.H{
			"msg": "参数错误",
		},
	})
}