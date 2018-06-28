// 外部应用认证模块
// 外部应用调用该模块提的接口验证访问用户是否具有相关操作的权限
package openAuth

import (
	"time"
	"fmt"
	"net/http"
	"encoding/json"
	"portal/model"
	"portal/util"
	"github.com/gin-gonic/gin"
)
type Query struct {
	TypeId int `json:"typeid" binding:"required,min:1"`
}
// header: token
// params: url?query={typeid:1}
// typeid 表示应用权限验证类型
// typeid = 1 无验证
// typeid = 2 token验证
// typeid = 3 接口权限验证
func Auth(c *gin.Context) {
	var params Query
	if err := json.Unmarshal([]byte(c.Query("query")), &params); err != nil {
		util.RespondBadRequest(c)
		return
	}
	// no auth
	if params.TypeId == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"error": gin.H{
				"msg": "",
			},
		})
		return
	}
	// token
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 1,
			"error": gin.H{
				"msg": "token不能为空",
			},
		})
		return
	}
	// 记录本次请求
  log := model.Log{
		Url:    c.Request.URL.Path, 
		Method: c.Request.Method, 
		Proto:  c.Request.Proto, 
		Agent:  c.Request.UserAgent(), 
		Host:   c.Request.Host,
		CreateAt: time.Now().Format(util.TimeFormat),
	}

	fmt.Println(log)
	// token := c.Request.Header.Get("token")
  
}
