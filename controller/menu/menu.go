package menu

import (
	"fmt"
	"net/http"

	"portal/util"
	"portal/model"
	"portal/service"

	"github.com/gin-gonic/gin"
)
func CreateRouter(c *gin.Context) {
	var jsonBody model.Route

	err := c.BindJSON(&jsonBody)
	if err != nil {
		fmt.Println(err)
		util.RespondBadRequest(c)
		return
	}
	code, msg := service.CreateRouter(jsonBody)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"error": gin.H{
			"msg": msg,
		},
	})
}