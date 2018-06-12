package util

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
// no signin
func RequireSignin(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 1,
		"error": gin.H{
			"msg": "未登录",
		},
	})
	c.Abort()
}
// admin required
func RequireAdmin(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 1,
		"error": gin.H{
			"msg": "没有访问权限",
		},
	})
	c.Abort()
}
// Response
type BaseResponse struct {
	Code  int   `json:"code"`
	Error struct {
		Msg interface{} `json:"msg"`
	} `json:"error"`
}