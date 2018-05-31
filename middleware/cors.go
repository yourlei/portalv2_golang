package middleware

import (
	"github.com/gin-gonic/gin"
)

//解决跨域问题，如前端链接不同，请改下面的地址
func Cors(c *gin.Context) {
	if "OPTIONS" == c.Request.Method {
		c.Status(200)
	}
	c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:8020")
	c.Header("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,PATCH,OPTIONS")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Content-Type", "application/json; charset=utf-8")
}
